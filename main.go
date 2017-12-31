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
    "os"
    "time"
)

func initialize() error {
    if val := os.Getenv(ENV_DEBUG); val == TRUE_STRVAL {
        DEBUG = true
    }
    if dir := os.Getenv(ENV_CACHE); dir != BLANK_LINE {
        CACHE_DIR = dir
    }
    return nil
}

func run(skyle *Skyle) {
    start := time.Now()
    if !DEBUG {
        defer Catch()
    }
    skyle.init().parse()
    skyle.run()
    skyle.save()
    opts := skyle.Profile().Program().Options()
    if opts.Timeit() {
        fmt.Printf("--\nRuntime %v\n", time.Since(start))
    }
}

func help(args *Args) {
    PrintLogo()
    args.Help()
}

func main() {
    if err := initialize(); err != nil {
        panic(err)
    }
    args := NewArgs()
    skyle := NewSkyle(args)
    if args.HasProfile() {
        run(skyle)
    } else {
        help(args)
    }
}
