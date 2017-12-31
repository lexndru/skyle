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
    "errors"
)

const (
    ENV_DEBUG         string  = "DEBUG"
    ENV_CACHE         string  = "CACHE"
    COMMENT_LINE      string  = "#"
    BLANK_LINE        string  = ""
    EMPTY_SPACE       string  = " "
    QUOTE             string  = `"`
    VAR_PREFIX        string  = "$"
    VAR_CURRENT       string  = "$$"
    VAR_PATTERN       string  = `\$([a-zA-Z][\w]+)`
    NON_CHARS_PATTERN string  = `\W`
    NON_EMPTY_STRING  string  = `[^\s]*`
    INSTR_MAP_LEN     int     = 1
    MIN_TOKENS        int     = 2
    TOKENS_LEN_TWO    int     = 2
    TOKENS_LEN_ONE    int     = 1
    FLAGS_SPLIT       string  = "="
    FLAGS_ARGS_INT    int     = 2
    CACHE_TIME        float64 = 1
    CACHE_INDEX       string  = "index"
    NOT_AVAILABLE     string  = "(n/a)"
    TRUE_STRVAL       string  = "true"
    WARNING_PREFIX    string  = "Warning: %s"
    FLUSH_ALL         string  = "all"
    MAX_ITER          uint16  = 65535
    ITER_STEP         uint16  = 1
    MODE_APPEND       string  = "append"
    MODE_WRITE        string  = "write"
    EXEC_ASYNC        string  = "async"
    EXEC_SYNC         string  = "sync"
    DEFAULT_TIMEOUT   int     = 30
    DEFAULT_SYNC      uint8   = 1
    INTERACTIVE       string  = "interactive"

    PROFILE string = "profile"
    TITLE   string = "title"
    PROBE   string = "probe"
    AGENT   string = "agent"
    FLAGS   string = "flags"
    OUTPUT  string = "output"
    FOLLOW  string = "follow"
    REMOVE  string = "remove"
    REPLACE string = "replace"
    PATTERN string = "pattern"
    GLUE    string = "glue"
    EXEC    string = "exec"
    SAVE    string = "save"
    KEEP    string = "keep"
    DUMP    string = "dump"
    NODE    string = "node"
    NEXT    string = "next"
    FLUSH   string = "flush"

    OPT_SYNC    string = "sync"
    OPT_PROXY   string = "proxy"
    OPT_CACHE   string = "cache"
    OPT_FORMAT  string = "format"
    OPT_MIME    string = "mime"
    OPT_TIMEIT  string = "timeit"
    OPT_VERBOSE string = "verbose"
    OPT_MAXITER string = "maxiter"
    OPT_TIMEOUT string = "timeout"
    OPT_MODE    string = "mode"
    OPT_EXEC    string = "exec"

    CSV_MIME           string = "text/csv"
    XML_MIME           string = "text/xml"
    HTML_MIME          string = "text/html"
    TEXT_MIME          string = "text/plain"
    JSON_MIME          string = "application/json"
    OCTECT_STREAM_MIME string = "application/octet-stream"

    HTTP_GET     string = "GET"
    USER_AGENT   string = "User-Agent"
    CONTENT_TYPE string = "Content-Type"

    FILE_PATTERN string = `^(?:\.{0,2}/)?(?:[\w\.]+/?)*$`
    URI_PATTERN  string = `^https?://[a-z0-9\-\.]+/.*?`

    PROFILE_MISSING          string = "Profile is missing: does the file exist?"
    DOCUMENT_MISSING         string = "Document is missing: does the file exist?"
    OUTPUT_MISSING           string = "Output is missing: use OUTPUT header to create an output file"
    OUTPUT_INVALID           string = "Output is invalid: use OUTPUT header to set a valid output filepath"
    FLAGS_INVALID            string = "One or more flags are invalid: cannot parse %s"
    UNSUPPORTED_PROTOCOL     string = "Unsupported protocol: expected local file or URI"
    UNSUPPORTED_FLAG         string = "Unsupported profile flag: %s"
    UNDEFINED_CONTEXT        string = "Undefined context: unsupported mime-type or unavailable resource"
    UNDEFINED_HANDLER        string = "Undefined handler: unexpected call to unavailable resource: %s"
    IGNORABLE_LINE_PASSED    string = "Cannot parse an empty line"
    INVALID_LINE_PASSED      string = "Cannot parse an invalid line: %s"
    INVALID_TOKEN_ARGS       string = "Invalid number of arguments on line: %s"
    UNEXPECTED_TOKEN         string = "Unexpected token in program: %s"
    UNEXPECTED_TOKEN_TAIL    string = "Unexpected header in instructions list: %s"
    UNEXPECTED_TOKEN_NEXT    string = "Unexpected NEXT function in instructions: NODE function missing"
    UNEXPECTED_INSTR_LEN     string = "Unexpected instruction length: must be equal to %d"
    BLANK_PROFILE_TITLE      string = "Profile header missing: TITLE is required"
    BLANK_PROFILE_PROBE      string = "Profile header missing: PROBE is required"
    REACHED_MAX_ITER_LIMIT   string = "MAX_ITER limit reached: %d"
    OPTION_VALUE_INVALID     string = "Header option %s has an invalid value"
    OPTION_VALUE_UNSUPPORTED string = "Header option %s has an unsupported value"
    UNSUPPORTED_HANDLER      string = "Unsupported handler"
    UNSUPPORTED_OUTPUT_MODE  string = "Unsupported output mode: %s"
    UNSUPPORTED_INPUT_MODE   string = "Unsupported mimetype: %s; fallback to text/plain"
    UNSUPPORTED_MIME_TYPE    string = "Unsupported mimetype: %s; please fix"
    UNSUPPORTED_EXEC_MODE    string = "Unsupported EXEC mode: %s; using default"
    INCOMPATIBLE_APPEND_FILE string = "Incompatible file to append: generating random file %s"

    HTTP_PROXY string = "HTTP_PROXY"
    SKYLE_UA   string = "Mozilla/5.0 (compatible; Skyle/1.0; +http://codeissues.net)"
)

var (
    SKYLE_VERSION, SKYLE_BUILD, SKYLE_OSARCH string

    CACHE_DIR = "/tmp/skyle"
    DEBUG     = false
    CHARS     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    SUPPORT   = []string{TEXT_MIME, HTML_MIME, XML_MIME, JSON_MIME, OCTECT_STREAM_MIME}

    UNPARSABLE_ERROR        error = errors.New("Unparsable error found")
    ERR_LOOP_NOT_EMPTY      error = errors.New("Unprocessed nodes await")
    PROFILE_NOT_INITIALIZED error = errors.New("Profile is not initialized")
)
