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
    "os"
    "io/ioutil"
    "bytes"
    "errors"
    "testing"
)

type Dummy struct {
    profile, probe, output              string
    profile_ctx, probe_ctx, output_ctx  []byte
}

func (d *Dummy) genericTest() error {
    if err := d.writeInput(d.profile, d.profile_ctx); err != nil {
        return err
    }
    if err := d.writeInput(d.probe, d.probe_ctx); err != nil {
        return err
    }
    if err := d.launchSkyle(); err != nil {
        return err
    }
    if err := d.checkOutput(); err != nil {
        return err
    }
    if err := d.cleanTmpFiles(d.profile, d.probe, d.output); err != nil {
        return err
    }
    return nil
}

func (d *Dummy) launchSkyle() error {
    args := &Args{
        profile: String{true, d.profile},
        probe:   String{true, d.probe},
        output:  String{true, d.output},
    }
    skyle := NewSkyle(args)
    skyle.init().parse()
    skyle.run()
    skyle.save()
    return nil
}

func (d *Dummy) writeInput(filename string, ctx []byte) error {
    return ioutil.WriteFile(filename, ctx, 0644)
}

func (d *Dummy) checkOutput() error {
    data, err := ioutil.ReadFile(d.output)
    if err != nil {
        return err
    }
    if bytes.Equal(d.output_ctx, data) {
        return nil
    }
    return errors.New("Different output context")
}

func (d *Dummy) cleanTmpFiles(files ...string) error {
    for _, fp := range files {
        if err := os.Remove(fp); err != nil {
            return err
        }
    }
    return nil
}

func TestMain(t *testing.T) {
    // noop
}
