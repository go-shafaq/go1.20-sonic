//go:build !go1.21
// +build !go1.21

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resolver

import (
    `github.com/go-shafaq/go1.20-encoding-json`
    `reflect`
    _ `unsafe`
)

type StdField struct {
    name        string
    nameBytes   []byte
    equalFold   func()
    nameNonEsc  string
    nameEscHTML string
    tag         bool
    index       []int
    typ         reflect.Type
    omitEmpty   bool
    quoted      bool
    encoder     func()
}

type StdStructFields struct {
    list      []StdField
    nameIndex map[string]int
}

func typeFields(t reflect.Type) StdStructFields {
	out := json.TypeFields(t)
	in := StdStructFields{}

	in.nameIndex = out.NameIndex
	in.list = make([]StdField, len(out.List))
	for i, f := range out.List {
		of := StdField{}

		of.name = f.Name
		of.nameBytes = f.NameBytes
		of.equalFold = nil
		of.nameNonEsc = f.NameNonEsc
		of.nameEscHTML = f.NameEscHTML
		of.tag = f.Tag
		of.index = f.Index
		of.typ = f.Typ
		of.omitEmpty = f.OmitEmpty
		of.quoted = f.Quoted
		of.encoder = nil

		in.list[i] = of
	}

	return in
}
