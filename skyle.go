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
    "fmt"
    "time"
)

type Source struct {
    profile, probe string
}

type Skyle struct {
    args     *Args
    env      *Option
    src      *Source
    profile  *Profile
    document *Document
}

func NewSkyle(a *Args) *Skyle {
    s := &Skyle{args: a}
    s.src = &Source{}
    s.env = &Option{}
    return s
}

func (s *Skyle) init() *Skyle {
    if val, ok := s.args.Profile(); !ok || val == BLANK_LINE {
        ParseError(PROFILE_MISSING)
    } else {
        s.src.profile = val
    }
    return s
}

func (s *Skyle) setup() *Skyle {
    if s.profile == nil {
        ParseError(PROFILE_NOT_INITIALIZED)
    }
    s.env = s.profile.options
    s.environment()
    return s
}

func (s *Skyle) environment() {
    if val, ok := s.args.Format(); ok {
        s.profile.options.format = val
    }
    if val, ok := s.args.Mime(); ok {
        s.profile.options.mime = val
    }
    if val, ok := s.args.Proxy(); ok {
        s.profile.options.proxy = val
        SetHTTPProxy(val)
    }
    if val, ok := s.args.Mode(); ok {
        s.profile.options.mode = val
    }
    if val, ok := s.args.Exec(); ok {
        s.profile.options.exec = val
    }
    if val, ok := s.args.Cache(); ok {
        s.profile.options.cache = val
    }
    if val, ok := s.args.Verbose(); ok {
        s.profile.options.verbose = val
    }
    if val, ok := s.args.Timeit(); ok {
        s.profile.options.timeit = val
    }
    if val, ok := s.args.Sync(); ok {
        s.profile.options.sync = uint8(val)
    }
    if val, ok := s.args.MaxIter(); ok {
        s.profile.options.maxiter = uint16(val)
    }
    if val, ok := s.args.Timeout(); ok {
        s.profile.options.timeout = time.Duration(val)
    }
}

func (s *Skyle) parse() *Skyle {
    s.profile = NewProfile(s)
    s.profile.init()
    s.profile.setup()
    s.setup()
    if val, ok := s.args.Output(); ok && val != BLANK_LINE {
        s.profile.program.header.output = val
    }
    if val, ok := s.args.Agent(); ok && val != BLANK_LINE {
        s.profile.program.header.agent = val
    }
    if val, ok := s.args.Probe(); ok && val != BLANK_LINE {
        s.profile.program.header.probe = val
        s.src.probe = val
    }
    if s.src.probe == BLANK_LINE {
        if s.profile.program.header.probe == BLANK_LINE {
            ParseError(DOCUMENT_MISSING)
        }
        s.src.probe = s.profile.program.header.probe
    }
    s.document = NewDocument(s)
    s.document.init()
    return s
}

func (s *Skyle) run() *Skyle {
    ctx := s.DocumentContextSlice()
    s.profile.program.setTmpValues(ctx)
    s.profile.run()
    return s
}

func (s *Skyle) save() *Skyle {
    out := NewOutput(s.profile.program)
    out.Write()
    return s
}

func (s *Skyle) Profile() *Profile {
    return s.profile
}

func (s *Skyle) Document() *Document {
    return s.document
}

func (s *Skyle) DocumentContextSlice() []string {
    ctx := s.document.String()
    return []string{ctx}
}

func PrintLogo() {
    fmt.Println(`     _          _      `)
    fmt.Println(` ___| | ___   _| | ___ `)
    fmt.Println(`/ __| |/ / | | | |/ _ \`)
    fmt.Println(`\__ \   <| |_| | |  __/`)
    fmt.Println(`|___/_|\_\\__, |_|\___|`)
    fmt.Println(`          |___/        `)
    fmt.Println()
    fmt.Println("Skyle (c) alex@codeissues.net")
    fmt.Println()
    fmt.Println("Version:", SKYLE_VERSION)
    fmt.Println("  Build:", SKYLE_BUILD)
    fmt.Println("OS/Arch:", SKYLE_OSARCH)
    fmt.Println()
}
