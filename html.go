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
    "bytes"
    "fmt"
    "strings"

    "github.com/moovweb/gokogiri"
    "github.com/moovweb/gokogiri/html"
    xhtml "golang.org/x/net/html"
)

type HtmlDocument struct {
    XmlDocument
    RootNode *html.HtmlDocument
}

func (hd *HtmlDocument) cleanup() {
    hd.RootNode.Free()
}

func (hd *HtmlDocument) initialize(pr *Program) {
    hd.TextDocument.program = pr
    doc := pr.profile.skyle.Document().String()
    hd.setGlobalRootNode(doc)
}

func (hd *HtmlDocument) setGlobalRootNode(doc string) {
    reader := strings.NewReader(doc)
    res, err := xhtml.Parse(reader)
    if err != nil {
        hd.program.LogError(err)
    }
    buf := &bytes.Buffer{}
    xhtml.Render(buf, res)
    hd.RootNode, err = gokogiri.ParseHtml(buf.Bytes())
    if err != nil {
        hd.program.LogError(err)
    }
}

func (hd *HtmlDocument) setGlobalNodeList(path string) {
    if hd.RootNodeChanged {
        hd.NodeList, _ = hd.LastRootNode.Search(path)
    } else {
        hd.NodeList, _ = hd.RootNode.Search(path)
    }
}

func (hd *HtmlDocument) RunFollow(s *Skyle, args []string) error {
    path := strings.Join(args, EMPTY_SPACE)
    s.Profile().Program().LogMessage("Following path %s", path)
    hd.setGlobalNodeList(path)
    lastFollow := []string{}
    for _, hd.LastNode = range hd.NodeList {
        content := hd.LastNode.Content()
        if strings.HasSuffix(path, "text()") && content == BLANK_LINE {
            s.Profile().Program().LogMessage("Warning: no text content found")
        }
        lastFollow = append(lastFollow, content)
    }
    s.Profile().Program().LogMessage(fmt.Sprintf("Found %d item(s)", len(hd.LastFollow)))
    s.Profile().Program().Store(&hd.LastFollow, lastFollow)
    return nil
}

func (hd *HtmlDocument) RunNode(s *Skyle, args []string) error {
    if hd.RootNodeList == nil {
        path := strings.Join(args, EMPTY_SPACE)
        hd.setGlobalNodeList(path)
        hd.RootNodeList = hd.NodeList
        if len(hd.RootNodeList) > hd.NodeListCounter {
            hd.LastRootNode = hd.RootNodeList[hd.NodeListCounter]
            hd.RootNodeChanged = true
            hd.NodeListCounter += 1
        }
    }
    return nil
}
