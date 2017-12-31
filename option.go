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
    "time"
    "strconv"
)

type Option struct {
    cache, verbose, timeit          bool
    format, proxy, mode, exec, mime string
    sync                            uint8
    maxiter                         uint16
    timeout                         time.Duration
}

func NewOptions() *Option {
    f := &Option{
        cache:   false,
        verbose: false,
        timeit:  false,
        format:  CSV_MIME,
        mime:    HTML_MIME,
        mode:    MODE_APPEND,
        exec:    EXEC_SYNC,
        sync:    DEFAULT_SYNC,
        timeout: time.Duration(DEFAULT_TIMEOUT),
        maxiter: MAX_ITER,
    }
    return f
}

func (o *Option) setCache(s string) {
    if val, err := strconv.ParseBool(s); err == nil {
        o.cache = val
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_CACHE)
    }
}

func (o *Option) setVerbose(s string) {
    if val, err := strconv.ParseBool(s); err == nil {
        o.verbose = val
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_VERBOSE)
    }
}

func (o *Option) setTimeit(s string) {
    if val, err := strconv.ParseBool(s); err == nil {
        o.timeit = val
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_TIMEIT)
    }
}

func (o *Option) setSync(s string) {
    if val, err := strconv.Atoi(s); err == nil {
        o.sync = uint8(val)
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_SYNC)
    }
}

func (o *Option) setTimeout(s string) {
    if val, err := strconv.Atoi(s); err == nil {
        o.timeout = time.Duration(val)
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_TIMEOUT)
    }
}

func (o *Option) setMaxIter(s string) {
    if val, err := strconv.ParseUint(s, 10, 16); err == nil {
        o.maxiter = uint16(val)
    } else {
        ParseWarning(OPTION_VALUE_INVALID, OPT_MAXITER)
    }
}

func (o *Option) setFormat(s string) {
    switch s {
    case CSV_MIME:
        o.format = s
    default:
        ParseError(OPTION_VALUE_UNSUPPORTED, OPT_FORMAT)
    }
}

func (o *Option) setMime(s string) {
    switch s {
    case HTML_MIME:
        o.mime = s
    case XML_MIME:
        o.mime = s
    case TEXT_MIME:
        o.mime = s
    default:
        ParseError(UNSUPPORTED_MIME_TYPE, s)
    }
}

func (o *Option) setMode(s string) {
    switch s {
    case MODE_APPEND, MODE_WRITE:
        o.mode = s
    default:
        ParseError(UNSUPPORTED_OUTPUT_MODE, s)
    }
}

func (o *Option) setProxy(s string) {
    if s == BLANK_LINE {
        ParseWarning(OPTION_VALUE_UNSUPPORTED, OPT_PROXY)
    } else {
        o.proxy = s
    }
}

func (o *Option) setExecMode(s string) {
    if s == EXEC_SYNC {
        o.exec = EXEC_SYNC
    } else if s == EXEC_ASYNC {
        o.exec = EXEC_ASYNC
    } else {
        ParseError(UNSUPPORTED_EXEC_MODE, s)
    }
}

func (o *Option) Sync() uint8 {
    return o.sync
}

func (o *Option) Timeout() time.Duration {
    return o.timeout
}

func (o *Option) Cache() bool {
    return o.cache
}

func (o *Option) Verbose() bool {
    return o.verbose
}

func (o *Option) Timeit() bool {
    return o.timeit
}

func (o *Option) Format() string {
    return o.format
}

func (o *Option) Mime() string {
    return o.mime
}

func (o *Option) Mode() string {
    return o.mode
}

func (o *Option) ExecAsyncMode() bool {
    return o.exec == EXEC_ASYNC
}
