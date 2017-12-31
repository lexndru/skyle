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
    "bufio"
    "bytes"
    "strconv"
    "strings"
)

type Profile struct {
    skyle   *Skyle
    token   *Token
    program *Program
    context []byte
    options *Option
    parser  Parser
}

func NewProfile(s *Skyle) *Profile {
    p := &Profile{skyle: s}
    p.options = NewOptions()
    p.token = NewToken()
    p.context, _ = ReadFromFile(s.src.profile)
    p.program = NewProgram()
    p.program.setProfile(p)
    return p
}

func (p *Profile) init() {
    reader := bytes.NewReader(p.context)
    scanner := bufio.NewScanner(reader)
    instructions := false
    for scanner.Scan() {
        line := scanner.Text()
        tokens := strings.Fields(line)
        if strings.HasPrefix(line, COMMENT_LINE) || len(tokens) == 0 {
            continue
        } else if len(tokens) < MIN_TOKENS {
            ParseWarning(INVALID_TOKEN_ARGS, line)
            continue
        }
        fun, args := tokens[0], tokens[1:]
        if p.token.isHead(fun) {
            if instructions {
                ParseError(UNEXPECTED_TOKEN_TAIL, fun)
            }
            p.program.setHeader(fun, args)
        } else if p.token.isTail(fun) {
            if !instructions {
                instructions = true
            }
            p.program.setFunction(fun, args)
        } else {
            ParseError(UNEXPECTED_TOKEN, fun)
        }
    }
}

func (p *Profile) setup() *Profile {
    for k, v := range p.program.header.Flags() {
        switch k {
        case OPT_CACHE:
            p.options.setCache(v)
        case OPT_VERBOSE:
            p.options.setVerbose(v)
        case OPT_TIMEIT:
            p.options.setTimeit(v)
        case OPT_SYNC:
            p.options.setSync(v)
        case OPT_TIMEOUT:
            p.options.setTimeout(v)
        case OPT_MAXITER:
            p.options.setMaxIter(v)
        case OPT_MODE:
            p.options.setMode(v)
        case OPT_MIME:
            p.options.setMime(v)
        case OPT_FORMAT:
            p.options.setFormat(v)
        case OPT_PROXY:
            p.options.setProxy(v)
        case OPT_EXEC:
            p.options.setExecMode(v)
        default:
            ParseWarning(UNSUPPORTED_FLAG, k)
        }
    }
    return p
}

func (p *Profile) run() {
    p.parser = NewParser(p.options.mime)
    p.parser.initialize(p.program)
    p.program.registerParser(p.parser)
    p.program.readInstructions().loopInstructions()
    p.parser.cleanup()
}

func (p *Profile) MaxIterLimit() uint16 {
    if p.options.maxiter > 0 {
        return p.options.maxiter
    }
    return MAX_ITER
}

func (p *Profile) Options() *Option {
    return p.options
}

func (p *Profile) MaxIterLimitString() string {
    return strconv.FormatUint(uint64(p.MaxIterLimit()), 10)
}

func (p *Profile) String() string {
    return string(p.context)
}

func (p *Profile) Token() *Token {
    return p.token
}

func (p *Profile) Program() *Program {
    return p.program
}
