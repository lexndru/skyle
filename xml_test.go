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
)

func TestXMLDocument(t *testing.T) {
    pfctx := []byte(`
follow //subject/@value
save subject
follow //subject/question/text()
save question
follow //subject/answer/text()
save answer
`)
    pbctx := []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<interview>
    <subject value="Introduction">
        <question>What's your name?</question>
        <answer>Skyle.</answer>
    </subject>
    <subject value="Introduction">
        <question>How are you?</question>
        <answer>I'm fine, thanks.</answer>
    </subject>
    <subject value="Interview">
        <question>What do you think of this interview?</question>
        <answer>It's nice!</answer>
    </subject>
</interview>
`)
    otctx := []byte(`subject,question,answer
Introduction,What's your name?,Skyle.
Introduction,How are you?,"I'm fine, thanks."
Interview,What do you think of this interview?,It's nice!
`)
    d := &Dummy{
        profile:     "/tmp/skyle_test_xmldoc.sky",
        probe:       "/tmp/skyle_test_xmldoc.txt",
        output:      "/tmp/skyle_test_xmldoc.csv",
        profile_ctx: pfctx,
        probe_ctx:   pbctx,
        output_ctx:  otctx,
    }
    if err := d.genericTest(); err != nil {
        t.Fail()
    }
}
