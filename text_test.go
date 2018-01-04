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

func TestInterview(t *testing.T) {
    pfctx := []byte(`
pattern Q:\s+(.*)\n
save question
flush all
pattern A:\s+(.*)\n
save answer
`)
    pbctx := []byte(`
Q: What's your name?
A: Skyle.
Q: How are you?
A: I'm fine, thanks.
Q: What do you think of this interview?
A: It's nice!
`)
    otctx := []byte(`question,answer
What's your name?,Skyle.
How are you?,"I'm fine, thanks."
What do you think of this interview?,It's nice!
`)
    d := &Dummy{
        profile:     "/tmp/skyle_test_interview.sky",
        probe:       "/tmp/skyle_test_interview.txt",
        output:      "/tmp/skyle_test_interview.csv",
        profile_ctx: pfctx,
        probe_ctx:   pbctx,
        output_ctx:  otctx,
    }
    if err := d.genericTest(); err != nil {
        t.Fail()
    }
}
