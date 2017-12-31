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
)

type Document struct {
    skyle    *Skyle
    resource string
    mimetype string
    context  []byte
}

type File struct {
    doc      *Document
    filepath string
    context  []byte
}

type URI struct {
    doc       *Document
    link      string
    userAgent string
    timeout   time.Duration
    context   []byte
    cached    bool
}

type Reader interface {
    Read() ([]byte, string, bool)
    Cache() (string, error)
    IsCached() bool
}

func NewFile(d *Document) *File {
    f := &File{}
    f.doc = d
    f.filepath = d.resource
    f.context = make([]byte, 0)
    return f
}

func (f *File) Read() ([]byte, string, bool) {
    body, mime := ReadFromFile(f.filepath)
    f.context = body
    if format, ok := HasSupport(mime); ok {
        return body, format, ok
    }
    return body, BLANK_LINE, false
}

func (f *File) Cache() (string, error) {
    return f.filepath, nil
}

func (f *File) IsCached() bool {
    return true
}

func NewURI(d *Document) *URI {
    u := &URI{}
    u.doc = d
    u.link = d.resource
    u.userAgent = d.skyle.profile.program.header.Agent()
    u.timeout = d.skyle.profile.options.Timeout()
    u.context = []byte{}
    return u
}

func (u *URI) Read() ([]byte, string, bool) {
    if fp, ok, _ := CacheExists(u); ok {
        body, mime := ReadFromFile(*fp)
        if format, ok := HasSupport(mime); ok {
            u.cached = len(body) > 0
            return body, format, ok
        }
    }
    res, err := HttpRequest(u)
    if err != nil {
        ParseError(err)
    }
    mime := res.Header.Get(CONTENT_TYPE)
    if _, ok := HasSupport(mime); !ok {
        return []byte{}, mime, ok
    }
    u.cached = false
    u.context = ReadFromHttpTransfer(res)
    return u.context, mime, true
}

func (u *URI) IsCached() bool {
    return u.cached
}

func (u *URI) Cache() (string, error) {
    fp, err := CacheURLToFile(u)
    return fp, err
}

func NewDocument(s *Skyle) *Document {
    d := &Document{skyle: s}
    d.resource = s.src.probe
    d.mimetype = BLANK_LINE
    d.context = []byte{}
    return d
}

func (d *Document) String() string {
    return string(d.context)
}

func (d *Document) init() {
    var reader Reader
    if d.isFile() {
        reader = NewFile(d)
    } else if d.isURI() {
        reader = NewURI(d)
    } else {
        ParseError(UNSUPPORTED_PROTOCOL)
    }
    if ok := d.setContext(reader); !ok {
        ParseError(UNDEFINED_CONTEXT)
    }
}

func (d *Document) isFile() bool {
    return StringMatchesCount(d.resource, FILE_PATTERN) > 0
}

func (d *Document) isURI() bool {
    return StringMatchesCount(d.resource, URI_PATTERN) > 0
}

func (d *Document) setContext(r Reader) bool {
    context, mime, ok := r.Read()
    if !ok {
        return false
    }
    d.context = context
    d.mimetype = mime
    if !r.IsCached() {
        if _, err := r.Cache(); err != nil {
            ParseError(err)
        }
    }
    return true
}

func (d *Document) Context() []byte {
    return d.context
}

func (d *Document) isHTML() bool {
    return StringContainsString(d.mimetype, HTML_MIME)
}

func (d *Document) isXML() bool {
    return StringContainsString(d.mimetype, XML_MIME)
}
