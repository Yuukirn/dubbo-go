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

package main

import (
	"flag"
)

import (
	"dubbo.apache.org/dubbo-go/v3/cmd/dubbo-java2proto/generator"
	"dubbo.apache.org/dubbo-go/v3/cmd/dubbo-java2proto/parser"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "file", "", "the path to the java file to be parsed")

	flag.Parse()

	jp, err := parser.NewParser(filepath)
	if err != nil {
		panic(err)
	}
	jp.ParseFile()

	gen := generator.NewGenerator(filepath + ".proto")
	err = gen.GenProto(jp)
	if err != nil {
		panic(err)
	}
}
