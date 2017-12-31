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

type Token struct {
    head []string
    tail []string
}

func NewToken() *Token {
    t := &Token{}
    t.head = []string{TITLE, PROBE, AGENT, FLAGS, OUTPUT}
    t.tail = []string{FOLLOW, REMOVE, REPLACE, PATTERN, GLUE, EXEC, SAVE, KEEP, DUMP, FLUSH, NODE, NEXT}
    return t
}

func (t *Token) is(list []string) func(string) bool {
    return func(str string) bool {
        for _, tok := range list {
            if str == tok {
                return true
            }
        }
        return false
    }
}

func (t *Token) isHead(str string) bool {
    return t.is(t.head)(str)
}

func (t *Token) isTail(str string) bool {
    return t.is(t.tail)(str)
}
