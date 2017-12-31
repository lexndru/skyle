// Copyright 2017 Alexandru Catrina
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
    "encoding/csv"
    "os"
)

type Output struct {
    program   *Program
    filepath  string
    generated bool
    data      map[string][]string
    headers   []string
    records   [][]string
    rows      int
}

func NewOutput(p *Program) *Output {
    o := &Output{program: p}
    o.init()
    o.parse()
    return o
}

func (o *Output) init() *Output {
    if fp, ok := o.program.Header().Output(); ok {
        o.filepath = fp
    } else {
        o.filepath = GenerateTmpFilename()
        o.generated = true
    }
    o.headers = o.program.Output()
    o.data = o.program.Cache()
    for _, v := range o.data {
        if len(v) > o.rows {
            o.rows = len(v)
        }
    }
    o.records = make([][]string, o.rows+1)
    o.records[0] = o.headers
    return o
}

func (o *Output) parse() *Output {
    for i := 1; i <= o.rows; i++ {
        row := make([]string, len(o.headers))
        for n, key := range o.headers {
            list := o.data[key]
            if index := i - 1; len(list) > index {
                row[n] = list[index]
            } else {
                row[n] = BLANK_LINE
            }
        }
        o.records[i] = row
    }
    return o
}

func (o *Output) isCompatible() bool {
    if o.generated {
        return false
    }
    file, err := os.Open(o.filepath)
    if os.IsNotExist(err) {
        return true
    } else if err != nil {
        ParseError(err)
    }
    defer file.Close()
    r := csv.NewReader(file)
    if data, err := r.Read(); err != nil {
        ParseError(err)
    } else {
        if data == nil || o.headers == nil {
            return false
        }
        if len(data) != len(o.headers) {
            return false
        }
        for i := range data {
            if data[i] != o.headers[i] {
                return false
            }
        }
        if len(o.records) > 1 {
            o.records = o.records[1:]
        }
        return true
    }
    return false
}

func (o *Output) Write() string {
    var file *os.File
    var err error
    if o.program.Options().Mode() == MODE_APPEND {
        if o.isCompatible() {
            file, err = os.OpenFile(o.filepath, os.O_APPEND|os.O_WRONLY, 0644)
            if os.IsNotExist(err) {
                file, err = os.Create(o.filepath)
            }
        } else {
            o.filepath = GenerateTmpFilename()
            file, err = os.Create(o.filepath)
            ParseWarning(INCOMPATIBLE_APPEND_FILE, o.filepath)
        }
    } else if o.program.Options().Mode() == MODE_WRITE {
        file, err = os.Create(o.filepath)
    } else {
        ParseError(UNSUPPORTED_OUTPUT_MODE, o.program.Options().Mode())
    }
    if err != nil {
        ParseError(err)
    }
    defer file.Close()
    w := csv.NewWriter(file)
    w.WriteAll(o.records)
    if err = w.Error(); err != nil {
        ParseError(err)
    }
    return o.filepath
}
