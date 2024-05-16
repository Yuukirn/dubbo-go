/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
	"strconv"
)

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

import (
	"dubbo.apache.org/dubbo-go/v3/proto/hessian2_extend"
)

func GetGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string) {
	goType, pointer := FieldGoType(g, field)
	if pointer {
		goType = "*" + goType
	}
	return
}

// below is helper func that copy from protobuf-gen-go

// FieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func FieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	var isWrapper bool
	opt, ok := proto.GetExtension(field.Desc.Options(), hessian2_extend.E_FieldExtend).(*hessian2_extend.Hessian2FieldOptions)
	if ok && opt != nil {
		isWrapper = opt.IsWrapper
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
		if isWrapper {
			goType = "*" + goType
		}
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
		if isWrapper {
			goType = "*" + goType
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
		if isWrapper {
			goType = "*" + goType
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
		if isWrapper {
			goType = "*" + goType
		}
	case protoreflect.DoubleKind:
		goType = "float64"
		if isWrapper {
			goType = "*" + goType
		}
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := FieldGoType(g, field.Message.Fields[0])
		valType, _ := FieldGoType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

// TODO(Yuukirn): add case for wrapper
func FieldDefaultValue(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message, field *protogen.Field) string {
	opts, ok := proto.GetExtension(field.Desc.Options(), hessian2_extend.E_FieldExtend).(hessian2_extend.Hessian2FieldOptions)
	if ok && opts.IsWrapper {
		return "nil"
	}
	if field.Desc.IsList() {
		return "nil"
	}
	if field.Desc.HasDefault() {
		defVarName := "Default_" + m.GoIdent.GoName + "_" + field.GoName
		if field.Desc.Kind() == protoreflect.BytesKind {
			return "append([]byte(nil), " + defVarName + "...)"
		}
		return defVarName
	}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.StringKind:
		return `""`
	case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.BytesKind:
		return "nil"
	case protoreflect.EnumKind:
		val := field.Enum.Values[0]
		if val.GoIdent.GoImportPath == f.GoImportPath {
			return g.QualifiedGoIdent(val.GoIdent)
		} else {
			// If the enum value is declared in a different Go package,
			// reference it by number since the name may not be correct.
			// See https://github.com/golang/protobuf/issues/513.
			return g.QualifiedGoIdent(field.Enum.GoIdent) + "(" + strconv.FormatInt(int64(val.Desc.Number()), 10) + ")"
		}
	default:
		return "0"
	}
}
