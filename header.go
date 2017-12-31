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
)

type Header struct {
    title, probe, agent, flags, output string
}

func NewHeader() *Header {
    h := &Header{}
    return h
}

func (h *Header) Title() string {
    title := strings.Fields(h.title)
    if len(title) == 0 {
        ParseError(BLANK_PROFILE_TITLE)
    }
    return h.title
}

func (h *Header) Probe() string {
    probe := strings.Fields(h.probe)
    if len(probe) == 0 {
        ParseError(BLANK_PROFILE_PROBE)
    }
    return h.probe
}

func (h *Header) Agent() string {
    agent := strings.Fields(h.agent)
    if len(agent) == 0 {
        return SKYLE_UA
    }
    return h.agent
}

func (h *Header) Flags() map[string]string {
    flags := strings.Fields(h.flags)
    if len(flags) == 0 {
        return map[string]string{}
    }
    options := map[string]string{}
    for _, val := range flags {
        list := strings.SplitN(val, FLAGS_SPLIT, FLAGS_ARGS_INT)
        if len(list) == 1 {
            options[list[0]] = TRUE_STRVAL
        } else if len(list) >= 2 {
            options[list[0]] = strings.Join(list[1:], EMPTY_SPACE)
        } else {
            ParseError(FLAGS_INVALID, val)
        }
    }
    return options
}

func (h *Header) Output() (string, bool) {
    output := strings.Fields(h.output)
    if len(output) == 0 {
        return BLANK_LINE, false
    }
    return h.output, true
}
