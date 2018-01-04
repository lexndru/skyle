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
    "testing"
    "fmt"
)

func TestHTMLDocument(t *testing.T) {
    pfctx := []byte(`
node //div
follow ./h2/text()
save title
follow ./span/text()
remove \D
save views
follow ./p/text()
save content
follow ./a/@href
save file
next node
`)
    pbctx := []byte(`
<html>
    <head>
        <title>HTML Example</title>
    </head>
    <body>
        <div>
            <h2>An article</h2>
            <span>633 views</span>
            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
        </div>
        <div>
            <h2>Anoter article</h2>
            <span>542 views</span>
            <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
        </div>
        <div>
            <h2>Yet again an article</h2>
            <span>62 views</span>
            <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
        </div>
        <div>
            <h2>Still an article</h2>
            <a href="article.pdf">Download article</a>
        </div>
        <div>
            <h2>Last article</h2>
            <span>321 views</span>
        </div>
    </body>
</html>
    `)
    otctx := []byte(`title,views,content,file
An article,633,"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
Anoter article,542,"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
Yet again an article,62,Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.,
Still an article,,,article.pdf
Last article,321,,
`)
    d := &Dummy{
        profile:     "/tmp/skyle_test_htmldoc.sky",
        probe:       "/tmp/skyle_test_htmldoc.txt",
        output:      "/tmp/skyle_test_htmldoc.csv",
        profile_ctx: pfctx,
        probe_ctx:   pbctx,
        output_ctx:  otctx,
    }
    if err := d.genericTest(); err != nil {
        fmt.Println(err)
        t.Fail()
    }
}
