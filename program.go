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
    "strings"
    "strconv"
)

type Program struct {
    profile      *Profile
    memory       *Memory
    header       *Header
    tokens       []map[string][]string
    functions    map[string]*Function
    instructions [][2]int
    output       []string
    next         bool
    iteration    int
}

type Memory struct {
    temp  string
    pipe  []string
    cache map[string][]string
}

type Function struct {
    name string
    call func(*Skyle, []string) error
}

func NewProgram() *Program {
    p := &Program{next: false}
    p.header = NewHeader()
    p.memory = &Memory{
        cache: map[string][]string{},
    }
    p.functions = map[string]*Function{}
    p.output = []string{}
    return p
}

func (p *Program) registerParser(psr Parser) *Program {
    p.functions = map[string]*Function{
        FOLLOW:  &Function{FOLLOW, psr.RunFollow},
        REMOVE:  &Function{REMOVE, psr.RunRemove},
        PATTERN: &Function{PATTERN, psr.RunPattern},
        REPLACE: &Function{REPLACE, psr.RunReplace},
        GLUE:    &Function{GLUE, psr.RunGlue},
        KEEP:    &Function{KEEP, psr.RunKeep},
        SAVE:    &Function{SAVE, psr.RunSave},
        EXEC:    &Function{EXEC, psr.RunExec},
        DUMP:    &Function{DUMP, psr.RunDump},
        NODE:    &Function{NODE, psr.RunNode},
        NEXT:    &Function{NEXT, psr.RunNext},
        FLUSH:   &Function{FLUSH, psr.RunFlush},
    }
    return p
}

func (p *Program) setProfile(pr *Profile) *Program {
    p.profile = pr
    return p
}

func (p *Program) readInstructions() *Program {
    lastNodeIndex := -1
    for i, instr := range p.tokens {
        for k, _ := range instr {
            p.instructions = append(p.instructions, [2]int{i, i})
            if k == NODE {
                lastNodeIndex = i
            } else if k == NEXT {
                if lastNodeIndex > -1 {
                    list := p.instructions[:lastNodeIndex]
                    args := [2]int{lastNodeIndex, i}
                    p.instructions = append(list, args)
                } else {
                    ParseWarning(UNEXPECTED_TOKEN_NEXT)
                }
            }
        }
    }
    return p
}

func (p *Program) loopInstructions() *Program {
    list := make([]int, TOKENS_LEN_TWO)
    for _, i := range p.instructions {
        if len(i) != TOKENS_LEN_TWO {
            ParseError(UNEXPECTED_INSTR_LEN)
        } else {
            list[0] = i[0]
            list[1] = i[1] + 1
        }
        p.parse(list, ITER_STEP)
    }
    return p
}

func (p *Program) parse(list []int, limit uint16) *Program {
    if limit > p.profile.MaxIterLimit() {
        ParseError(REACHED_MAX_ITER_LIMIT, p.profile.MaxIterLimitString())
    } else if len(list) != TOKENS_LEN_TWO {
        ParseError(UNEXPECTED_INSTR_LEN, strconv.Itoa(TOKENS_LEN_TWO))
    }
    from, to := list[0], list[1]
    for _, i := range p.tokens[from:to] {
        if len(i) != INSTR_MAP_LEN {
            ParseError(UNEXPECTED_INSTR_LEN, strconv.Itoa(INSTR_MAP_LEN))
        }
        if err := p.call(i); err != nil {
            ParseError(err)
        }
    }
    if p.hasNext() {
        limit += ITER_STEP
        return p.parse(list, limit)
    }
    return p
}

func (p *Program) call(instr map[string][]string) error {
    for k, v := range instr {
        if fun, ok := p.functions[k]; ok {
            if err := fun.call(p.profile.skyle, v); err != nil {
                return err
            }
        } else {
            ParseWarning(UNEXPECTED_TOKEN, k)
        }
    }
    return nil
}

func (p *Program) setNext(val bool) {
    p.next = val
}

func (p *Program) hasNext() bool {
    return p.next
}

func (p *Program) setHeader(fun string, val []string) {
    switch v := strings.Join(val, EMPTY_SPACE); fun {
    case TITLE:
        p.header.title = v
    case PROBE:
        p.header.probe = v
    case AGENT:
        p.header.agent = v
    case FLAGS:
        p.header.flags = v
    case OUTPUT:
        p.header.output = v
    default:
        ParseError(UNEXPECTED_TOKEN, fun)
    }
}

func (p *Program) setFunction(fun string, val []string) {
    p.tokens = append(p.tokens, map[string][]string{fun: val})
    if fun == SAVE && len(val) > 0 {
        p.output = append(p.output, val[0])
        p.UpdateCache(val[0], []string{})
    }
}

func (p *Program) UpdateCache(key string, vals []string) {
    if cache, ok := p.memory.cache[key]; ok {
        p.memory.cache[key] = append(cache, vals...)
    } else {
        p.memory.cache[key] = vals
    }
}

func (p *Program) Cache() map[string][]string {
    return p.memory.cache
}

func (p *Program) Cached(key string) ([]string, bool) {
    if val, ok := p.memory.cache[key]; ok {
        return val, true
    }
    return []string{}, false
}

func (p *Program) Headers() *Header {
    return p.header
}

func (p *Program) Options() *Option {
    return p.profile.options
}
func (p *Program) Header() *Header {
    return p.header
}

func (p *Program) Output() []string {
    return p.output
}

func (p *Program) TmpValue() string {
    return p.memory.temp
}

func (p *Program) setTmpValue(tmp string) {
    p.memory.temp = tmp
}

func (p *Program) TmpValues() []string {
    return p.memory.pipe
}

func (p *Program) setTmpValues(tmp []string) {
    p.memory.pipe = tmp
}

func (p *Program) LogError(err error) {
    ParseError(err)
}

func (p *Program) LogMessage(text string, vals ...string) {
    if !p.Options().Verbose() {
        return
    }
    ParseMessage(text, vals...)
}

func (p *Program) Store(holder *[]string, list []string) {
    for i := uint8(len(list)); i < p.Options().Sync(); i++ {
        list = append(list, BLANK_LINE)
    }
    *holder = list
    p.setTmpValues(*holder)
}

func (p *Program) Template(args []string) []string {
    vars := p.Cache()
    results := []string{}
    replace := func(key string) string {
        str := Variable.FindStringSubmatch(key)
        if len(str) == 2 {
            if val, ok := vars[str[1]]; ok {
                v := strings.Join(val, EMPTY_SPACE)
                return strings.Replace(key, str[0], v, -1)
            }
        }
        return key
    }
    for _, key := range args {
        results = append(results, replace(key))
    }
    return results
}
