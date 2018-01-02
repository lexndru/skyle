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
    "io/ioutil"
    "os/exec"
    "os"
    "bytes"
    "testing"
)

func TestInterview(t *testing.T) {
    profile := "/tmp/skyle_test_interview.sky"
    profile_ctx := []byte(`
probe /tmp/skyle_test_interview.txt
flags mode=write
output /tmp/skyle_test_interview.csv
pattern Q:\s+(.*)\n
save question
flush all
pattern A:\s+(.*)\n
save answer
`)
    probe := "/tmp/skyle_test_interview.txt"
    probe_ctx := []byte(`
Q: What's your name?
A: Skyle.
Q: How are you?
A: I'm fine, thanks.
Q: What do you think of this interview?
A: It's nice!
`)
    output := "/tmp/skyle_test_interview.csv"
    output_ctx := []byte(`question,answer
What's your name?,Skyle.
How are you?,"I'm fine, thanks."
What do you think of this interview?,It's nice!
`)
    if err := ioutil.WriteFile(profile, profile_ctx, 0644); err != nil {
        t.Error(err)
    }
    if err := ioutil.WriteFile(probe, probe_ctx, 0644); err != nil {
        t.Error(err)
    }
    _, err := exec.Command(SKYLE_APP, profile).Output()
    if err != nil {
        t.Error(err)
    }
    data, err := ioutil.ReadFile(output)
    if err != nil {
        t.Error(err)
    }
    if !bytes.Equal(data, output_ctx) {
        t.Fail()
    }
    if err := os.Remove(profile); err != nil {
        t.Error(err)
    }
    if err := os.Remove(probe); err != nil {
        t.Error(err)
    }
    if err := os.Remove(output); err != nil {
        t.Error(err)
    }
}
