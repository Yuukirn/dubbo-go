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

package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

import (
	"dubbo.apache.org/dubbo-go/v3/cmd/protoc-gen-go-hessian2/util"
	"dubbo.apache.org/dubbo-go/v3/proto/hessian2_extend"
)

type Hessian2Go struct {
	*protogen.File

	Source       string
	ProtoPackage string
	Enums        []*Enum
	Messages     []*Message
}

type Enum struct {
	*protogen.Enum

	JavaClassName string
}

type Message struct {
	*protogen.Message

	JavaClassName string

	InnerMessages []*Message
	IsInheritance bool
	Fields        []*Field

	ExtendArgs bool
}

type Field struct {
	*protogen.Field

	Type         string
	DefaultValue string
}

func ProcessProtoFile(g *protogen.GeneratedFile, file *protogen.File) (*Hessian2Go, error) {
	hessian2Go := &Hessian2Go{
		File:         file,
		Source:       file.Proto.GetName(),
		ProtoPackage: file.Proto.GetPackage(),
	}

	for _, enum := range file.Enums {
		hessian2Go.Enums = append(hessian2Go.Enums, processProtoEnum(g, enum))
	}

	for _, message := range file.Messages {
		hessian2Go.Messages = append(hessian2Go.Messages, processProtoMessage(g, file, message))
	}

	return hessian2Go, nil
}

func processProtoEnum(g *protogen.GeneratedFile, e *protogen.Enum) *Enum {
	enum := &Enum{
		Enum: e,
	}
	g.QualifiedGoIdent(e.GoIdent)

	opts, ok := proto.GetExtension(e.Desc.Options(), hessian2_extend.E_EnumExtend).(*hessian2_extend.Hessian2EnumOptions)
	if !ok {
		return enum
	}
	enum.JavaClassName = opts.JavaClassName

	return enum
}

func processProtoMessage(g *protogen.GeneratedFile, file *protogen.File, m *protogen.Message) *Message {
	msg := &Message{
		Message: m,
	}

	if m.Desc.IsMapEntry() {
		return msg
	}

	for _, inner := range m.Messages {
		processedInnerMsg := processProtoMessage(g, file, inner)
		opts, ok := proto.GetExtension(inner.Desc.Options(), hessian2_extend.E_MessageExtend).(*hessian2_extend.Hessian2MessageOptions)
		if ok && opts != nil && opts.IsInheritance {
			processedInnerMsg.IsInheritance = true
		}
		msg.InnerMessages = append(msg.InnerMessages, processedInnerMsg)
	}

	var fields []*Field
	for _, field := range m.Fields {
		var typ string
		if field.Message != nil {
			opts, ok := proto.GetExtension(field.Message.Desc.Options(), hessian2_extend.E_MessageExtend).(*hessian2_extend.Hessian2MessageOptions)
			if ok && opts != nil && opts.ReferencePath != "" {
				typ = "*" + g.QualifiedGoIdent(protogen.GoIdent{
					GoName:       field.Message.GoIdent.GoName,
					GoImportPath: protogen.GoImportPath(opts.ReferencePath),
				})
			}
		}
		if typ == "" {
			typ = util.GetGoType(g, field)
		}
		fields = append(fields, &Field{
			Field:        field,
			Type:         typ,
			DefaultValue: util.FieldDefaultValue(g, file, m, field),
		})
	}
	msg.Fields = fields

	opts, ok := proto.GetExtension(m.Desc.Options(), hessian2_extend.E_MessageExtend).(*hessian2_extend.Hessian2MessageOptions)
	if !ok {
		panic(ErrNoMessageExtend)
	}
	if opts.JavaClassName == "" && !opts.ExtendArgs {
		panic(ErrNoJavaClassName)
	}
	msg.JavaClassName = opts.JavaClassName
	msg.ExtendArgs = opts.ExtendArgs

	return msg
}
