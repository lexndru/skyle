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
    "strings"

    "github.com/moovweb/gokogiri"
    "github.com/moovweb/gokogiri/xml"
)

type XmlDocument struct {
    TextDocument
    RootNodeChanged        bool
    NodeListCounter        int
    RootNode               *xml.XmlDocument
    LastNode, LastRootNode xml.Node
    NodeList, RootNodeList []xml.Node
}

func (xd *XmlDocument) cleanup() {
    xd.RootNode.Free()
}

func (xd *XmlDocument) initialize(pr *Program) {
    xd.TextDocument.program = pr
    doc := pr.profile.skyle.Document().String()
    xd.setGlobalRootNode(doc)
}

func (xd *XmlDocument) setGlobalRootNode(doc string) {
    var err error
    xd.RootNode, err = gokogiri.ParseXml([]byte(doc))
    if err != nil {
        xd.program.LogError(err)
    }
}

func (xd *XmlDocument) setGlobalNodeList(path string) {
    if xd.RootNodeChanged {
        xd.NodeList, _ = xd.LastRootNode.Search(path)
    } else {
        xd.NodeList, _ = xd.RootNode.Search(path)
    }
}

func (xd *XmlDocument) RunFollow(s *Skyle, args []string) error {
    path := strings.Join(args, EMPTY_SPACE)
    s.Profile().Program().LogMessage("Following path %s", path)
    xd.setGlobalNodeList(path)
    lastFollow := []string{}
    for _, xd.LastNode = range xd.NodeList {
        content := xd.LastNode.Content()
        if strings.HasSuffix(path, "text()") && content == BLANK_LINE {
            s.Profile().Program().LogMessage("Warning: no text content found")
        }
        lastFollow = append(lastFollow, content)
    }
    s.Profile().Program().LogMessage(fmt.Sprintf("Found %d item(s)", len(xd.LastFollow)))
    s.Profile().Program().Store(&xd.LastFollow, lastFollow)
    return nil
}

func (xd *XmlDocument) RunNode(s *Skyle, args []string) error {
    if xd.RootNodeList == nil {
        path := strings.Join(args, EMPTY_SPACE)
        xd.setGlobalNodeList(path)
        xd.RootNodeList = xd.NodeList
        if len(xd.RootNodeList) > xd.NodeListCounter {
            xd.LastRootNode = xd.RootNodeList[xd.NodeListCounter]
            xd.RootNodeChanged = true
            xd.NodeListCounter += 1
        }
    }
    return nil
}

func (xd *XmlDocument) RunNext(s *Skyle, args []string) error {
    if xd.RootNodeList == nil {
        return errors.New("NEXT function works better with NODE function")
    }
    if len(xd.RootNodeList) > xd.NodeListCounter {
        xd.LastRootNode = xd.RootNodeList[xd.NodeListCounter]
        xd.RootNodeChanged = true
        xd.NodeListCounter += 1
    } else {
        xd.RootNodeChanged = false
        xd.RootNodeList = nil
    }
    s.Profile().Program().setNext(xd.RootNodeChanged)
    return nil
}
