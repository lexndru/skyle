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
    "fmt"
    "os/exec"
    "regexp"
    "strings"
)

type TextDocument struct {
    program     *Program
    LastFollow  []string
    LastPattern []string
    LastReplace []string
    LastRemove  []string
    LastGlue    []string
    LastKeep    []string
    LastSave    []string
    LastExec    []string
}

func (td *TextDocument) cleanup() {
    // noop
}

func (td *TextDocument) initialize(pr *Program) {
    td.program = pr
}

func (td *TextDocument) RunFollow(s *Skyle, args []string) error {
    return errors.New("Cannot follow xpath for text document")
}

func (td *TextDocument) RunRemove(s *Skyle, args []string) error {
    values := s.Profile().Program().TmpValues()
    lastRemove := make([]string, len(values))
    if len(values) > 0 {
        vars := s.Profile().Program().Template(args)
        param := strings.Join(vars, EMPTY_SPACE)
        s.Profile().Program().LogMessage("Removing string:", param)
        for i, tmp := range values {
            lastRemove[i] = RemoveString(param, tmp)
        }
    }
    s.Profile().Program().Store(&td.LastRemove, lastRemove)
    return nil
}

func (td *TextDocument) RunPattern(s *Skyle, args []string) error {
    values := s.Profile().Program().TmpValues()
    lastPattern := []string{}
    if len(values) > 0 {
        param := strings.Join(args, EMPTY_SPACE)
        s.Profile().Program().LogMessage("Extracting pattern:", param)
        for _, tmp := range values {
            vals := ExtractPattern(param, tmp)
            lastPattern = append(lastPattern, vals...)
        }
    }
    s.Profile().Program().Store(&td.LastPattern, lastPattern)
    return nil
}

func (td *TextDocument) RunGlue(s *Skyle, args []string) error {
    values := s.Profile().Program().TmpValues()
    lastGlue := make([]string, len(values))
    if len(values) > 0 {
        vars := s.Profile().Program().Template(args)
        param := strings.Join(vars, EMPTY_SPACE)
        s.Profile().Program().LogMessage("Applying glue:", param)
        for i, tmp := range values {
            lastGlue[i] = Template(param, tmp)
        }
    }
    s.Profile().Program().Store(&td.LastGlue, lastGlue)
    return nil
}

func (td *TextDocument) RunKeep(s *Skyle, args []string) error {
    rule := NON_EMPTY_STRING
    if len(args) > 0 {
        rule = strings.Join(args, EMPTY_SPACE)
    }
    values := s.Profile().Program().TmpValues()
    lastKeep := []string{}
    if len(values) > 0 {
        re := regexp.MustCompile(rule)
        for _, v := range values {
            vals := strings.Fields(v)
            if len(vals) > 0 && re.MatchString(v) {
                lastKeep = append(lastKeep, v)
            }
        }
    }
    s.Profile().Program().Store(&td.LastKeep, lastKeep)
    return nil
}

func (td *TextDocument) RunSave(s *Skyle, args []string) error {
    if len(args) == 0 {
        return errors.New("Missing variable name to save")
    }
    key := args[0]
    values := s.Profile().Program().TmpValues()
    records := len(values)
    if len(args) == 1 {
        td.LastSave = values
    } else if len(args) >= 2 {
        val := strings.Join(args[1:], EMPTY_SPACE)
        td.LastSave = make([]string, records)
        td.LastSave[0] = val
        for i := 1; i < records; i++ {
            td.LastSave[i] = val
        }
    }
    msg := fmt.Sprintf("Saving %d item(s) as %s", records, strings.ToUpper(key))
    s.Profile().Program().LogMessage(msg)
    s.Profile().Program().UpdateCache(key, td.LastSave)
    return nil
}

func (td *TextDocument) RunExec(s *Skyle, args []string) error {
    vars := s.Profile().Program().Template(args)
    cmd := strings.Join(vars, EMPTY_SPACE)
    s.Profile().Program().LogMessage("Running command:", cmd)
    lastExec := []string{BLANK_LINE}
    stdout, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        s.Profile().Program().LogMessage(fmt.Sprintf("ERROR: %s", err.Error()))
    } else if len(stdout) > 0 {
        s.Profile().Program().LogMessage(fmt.Sprintf("OK: %s", stdout))
        lastExec[0] = string(stdout)
    } else {
        s.Profile().Program().LogMessage("OK: no output")
    }
    s.Profile().Program().Store(&td.LastExec, lastExec)
    return nil
}

func (td *TextDocument) RunFlush(s *Skyle, args []string) error {
    if len(args) == 0 {
        return errors.New("Missing value to flush")
    } else if len(args) == 1 && args[0] == FLUSH_ALL {
        s.Profile().Program().LogMessage("Flushing cache...")
        s.Profile().Program().setTmpValue(BLANK_LINE)
        s.Profile().Program().setTmpValues(s.DocumentContextSlice())
        s.Profile().Program().LogMessage("Cache is now restored to initial state")
    } else {
        s.Profile().Program().LogMessage("Unsupported value(s) to FLUSH at the moment")
    }
    return nil
}

func (td *TextDocument) RunReplace(s *Skyle, args []string) error {
    if len(args) < 2 {
        return errors.New("Not enough parameters to replace")
    }
    old, new_ := args[0], strings.Join(args[1:], EMPTY_SPACE)
    values := s.Profile().Program().TmpValues()
    lastReplace := make([]string, len(values))
    if len(values) > 0 {
        s.Profile().Program().LogMessage("Replacing:", old)
        for i, tmp := range values {
            lastReplace[i] = ReplaceString(old, tmp, new_)
        }
    }
    s.Profile().Program().Store(&td.LastReplace, lastReplace)
    return nil
}

func (td *TextDocument) RunDump(s *Skyle, args []string) error {
    msg := strings.Join(args, EMPTY_SPACE)
    s.Profile().Program().LogMessage("Memory dump:", msg)
    Dump(s.Profile().Program().TmpValues())
    return nil
}

func (td *TextDocument) RunNode(s *Skyle, args []string) error {
    return errors.New("Node is not available for text document")
}

func (td *TextDocument) RunNext(s *Skyle, args []string) error {
    return errors.New("Next is not available for text document")
}
