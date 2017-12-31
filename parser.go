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

type Parser interface {
    initialize(*Program)
    cleanup()
    RunFollow(*Skyle, []string) error
    RunRemove(*Skyle, []string) error
    RunPattern(*Skyle, []string) error
    RunReplace(*Skyle, []string) error
    RunGlue(*Skyle, []string) error
    RunKeep(*Skyle, []string) error
    RunSave(*Skyle, []string) error
    RunExec(*Skyle, []string) error
    RunDump(*Skyle, []string) error
    RunNode(*Skyle, []string) error
    RunNext(*Skyle, []string) error
    RunFlush(*Skyle, []string) error
}

func NewParser(mime string) Parser {
    var prs Parser
    if format, ok := HasSupport(mime); !ok {
        ParseWarning(UNSUPPORTED_MIME_TYPE, mime)
        prs = &TextDocument{}
    } else {
        switch format {
        case HTML_MIME:
            prs = &HtmlDocument{}
        case XML_MIME:
            prs = &XmlDocument{}
        case TEXT_MIME:
            prs = &TextDocument{}
        default:
            ParseError(UNSUPPORTED_INPUT_MODE, format)
        }
    }
    return prs
}
