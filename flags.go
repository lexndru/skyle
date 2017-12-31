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
    "flag"
    "fmt"
    "strconv"
    "strings"
)

type String struct {
    set   bool
    value string
}

type Boolean struct {
    set   bool
    value bool
}

type Number struct {
    set   bool
    value uint
}

type Args struct {
    profile, probe, output, agent   String
    format, proxy, mode, exec, mime String
    cache, verbose, timeit          Boolean
    sync, maxiter, timeout          Number
    repl                            bool
}

func (n *Number) Set(str string) error {
    if val, err := strconv.ParseUint(str, 10, 16); err != nil {
        return err
    } else {
        n.value = uint(val)
        n.set = true
    }
    return nil
}

func (n *Number) String() string {
    return fmt.Sprint(n.value)
}

func (b *Boolean) Set(str string) error {
    if val, err := strconv.ParseBool(str); err != nil {
        return err
    } else {
        b.value = val
        b.set = true
    }
    return nil
}

func (b *Boolean) String() string {
    return fmt.Sprint(b.value)
}

func (s *String) Set(str string) error {
    s.value = str
    s.set = true
    return nil
}

func (s *String) String() string {
    return s.value
}

func NewArgs() *Args {
    a := &Args{}
    a.init()
    if val, ok := a.support(); ok {
        a.profile.set = true
        a.profile.value = *val
    }
    return a
}

func (a *Args) HasProfile() bool {
    return a.profile.set && a.profile.value != BLANK_LINE
}

func (a *Args) Help() {
    flag.PrintDefaults()
}

func (a *Args) init() *Args {
    flag.Var(&a.profile, PROFILE, "Profile path of document")
    flag.Var(&a.output, OUTPUT, "Output path of results to save")
    flag.Var(&a.agent, AGENT, "Set user-agent to HTTP request")
    flag.Var(&a.probe, PROBE, "URL or path to probe document")
    flag.Var(&a.cache, OPT_CACHE, "Enable or disable cache")
    flag.Var(&a.verbose, OPT_VERBOSE, "Enable verbose output")
    flag.Var(&a.timeit, OPT_TIMEIT, "Measure execution time and display on exit")
    flag.Var(&a.mime, OPT_MIME, "Set input format (default text/html)")
    flag.Var(&a.format, OPT_FORMAT, "Set output format (default text/csv)")
    flag.Var(&a.proxy, OPT_PROXY, "HTTP proxy address (optional)")
    flag.Var(&a.mode, OPT_MODE, "Set file saving mode: append or write")
    flag.Var(&a.exec, OPT_EXEC, "Set shell exec mode: async or sync")
    flag.Var(&a.sync, OPT_SYNC, "Ensure all instructions return at least N values")
    flag.Var(&a.maxiter, OPT_MAXITER, "Change the maximum iteration limit")
    flag.Var(&a.timeout, OPT_TIMEOUT, "Change HTTP request timeout")
    flag.BoolVar(&a.repl, INTERACTIVE, false, "Launch interactive REPL")
    flag.Parse()
    return a
}

func (a *Args) support() (*string, bool) {
    if flag.NFlag() == 0 && flag.NArg() > 0 {
        list := make([]string, flag.NArg())
        for i := 0; i < flag.NArg(); i++ {
            list[i] = flag.Arg(i)
        }
        profile := strings.Join(list, EMPTY_SPACE)
        return &profile, true
    }
    return nil, false
}

func (a *Args) Profile() (string, bool) {
    r := a.profile
    return r.value, r.set
}

func (a *Args) Output() (string, bool) {
    r := a.output
    return r.value, r.set
}

func (a *Args) Agent() (string, bool) {
    r := a.agent
    return r.value, r.set
}

func (a *Args) Probe() (string, bool) {
    r := a.probe
    return r.value, r.set
}

func (a *Args) Cache() (bool, bool) {
    r := a.cache
    return r.value, r.set
}

func (a *Args) Verbose() (bool, bool) {
    r := a.verbose
    return r.value, r.set
}

func (a *Args) Timeit() (bool, bool) {
    r := a.timeit
    return r.value, r.set
}

func (a *Args) Format() (string, bool) {
    r := a.format
    return r.value, r.set
}

func (a *Args) Mime() (string, bool) {
    r := a.mime
    return r.value, r.set
}

func (a *Args) Proxy() (string, bool) {
    r := a.proxy
    return r.value, r.set
}

func (a *Args) Mode() (string, bool) {
    r := a.mode
    return r.value, r.set
}

func (a *Args) Exec() (string, bool) {
    r := a.exec
    return r.value, r.set
}

func (a *Args) Sync() (uint, bool) {
    r := a.sync
    return r.value, r.set
}

func (a *Args) MaxIter() (uint, bool) {
    r := a.maxiter
    return r.value, r.set
}

func (a *Args) Timeout() (uint, bool) {
    r := a.timeout
    return r.value, r.set
}
