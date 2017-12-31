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
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "regexp"
    "strings"
    "time"
)

var Variable = regexp.MustCompile(VAR_PATTERN)

func ReplaceString(param, str, val string) string {
    re := regexp.MustCompile(param)
    return re.ReplaceAllString(str, val)
}

func RemoveString(param, str string) string {
    return ReplaceString(param, str, BLANK_LINE)
}

func ExtractPattern(param, str string) (vals []string) {
    re := regexp.MustCompile(param)
    for _, res := range re.FindAllStringSubmatch(str, -1) {
        val := BLANK_LINE
        if len(res) == 2 {
            val = res[1]
        }
        vals = append(vals, val)
    }
    return vals
}

func Template(param, str string) string {
    return strings.Replace(param, VAR_CURRENT, str, -1)
}

func StringMatches(str, exp string) []string {
    regex := regexp.MustCompile(exp)
    return regex.FindStringSubmatch(str)
}

func StringMatchesCount(str, exp string) int {
    list := StringMatches(str, exp)
    return len(list)
}

func StringContainsString(str, sub string) bool {
    return strings.Contains(str, sub)
}

func GetMIMEType(b *[]byte) string {
    buffer, length := *b, 512
    return http.DetectContentType(buffer[:length])
}

func ReadFromHttpTransfer(r *http.Response) []byte {
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ParseError(err)
    }
    return body
}

func ReadFromFile(fp string) ([]byte, string) {
    buf, err := ioutil.ReadFile(fp)
    if err != nil {
        ParseError(err)
    }
    return buf, GetMIMEType(&buf)
}

func CacheExists(u *URI) (*string, bool, error) {
    if !u.doc.skyle.Profile().Options().Cache() {
        return nil, false, nil
    }
    uri, err := url.Parse(u.link)
    if err != nil {
        return nil, false, err
    }
    dir := fmt.Sprintf("%s/%s", CACHE_DIR, uri.Host)
    fp := CacheableFilepath(uri.Path, &dir)
    fi, err := os.Stat(fp)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, false, err
        }
        return &fp, true, err
    }
    if time.Since(fi.ModTime()).Hours() > CACHE_TIME {
        return &fp, false, nil
    }
    return &fp, true, nil
}

func CacheableFilepath(uri string, dir *string) string {
    if dir == nil {
        *dir = CACHE_DIR
    }
    re := regexp.MustCompile(NON_CHARS_PATTERN)
    fp := strings.TrimPrefix(re.ReplaceAllString(uri, "_"), "_")
    if fp == BLANK_LINE {
        fp = CACHE_INDEX
    }
    return fmt.Sprintf("%s/%s.tmp", *dir, fp)
}

func CacheURLToFile(u *URI) (string, error) {
    uri, err := url.Parse(u.link)
    if err != nil {
        return BLANK_LINE, err
    }
    os.Mkdir(CACHE_DIR, 0755)
    dir := fmt.Sprintf("%s/%s", CACHE_DIR, uri.Host)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.Mkdir(dir, 0755)
    }
    fp := CacheableFilepath(uri.Path, &dir)
    return fp, ioutil.WriteFile(fp, u.context, 0644)
}

func HttpRequest(u *URI) (*http.Response, error) {
    client := &http.Client{Timeout: time.Second * u.timeout}
    req, err := http.NewRequest(HTTP_GET, u.link, nil)
    if err != nil {
        ParseError(err)
    }
    req.Header.Set(USER_AGENT, u.userAgent)
    return client.Do(req)
}

func ParseError(err interface{}, args ...string) {
    var msg string
    switch err.(type) {
    case error:
        msg = err.(error).Error()
    case string:
        msg = err.(string)
    default:
        panic(UNPARSABLE_ERROR)
    }
    if DEBUG {
        if len(args) > 0 {
            msg = fmt.Sprintf(msg, args)
        }
        panic(msg)
    }
    if len(args) > 0 {
        log.Fatalf(msg, args)
    } else {
        log.Fatal(msg)
    }
}

func ParseWarning(err string, args ...string) {
    msg := fmt.Sprintf(err, strings.Join(args, EMPTY_SPACE))
    log.Printf(WARNING_PREFIX, msg)
}

func ParseMessage(msg string, args ...string) {
    if len(args) > 0 {
        log.Println(msg, strings.Join(args, EMPTY_SPACE))
    } else {
        log.Println(msg)
    }
}

func HasSupport(mime string) (string, bool) {
    for _, format := range SUPPORT {
        if StringContainsString(mime, format) {
            return format, true
        }
    }
    return BLANK_LINE, false
}

func GenerateTmpFilename() string {
    n := len(CHARS)
    buf := make([]byte, 16)
    for i, _ := range buf {
        buf[i] = CHARS[rand.Intn(n)]
    }
    return fmt.Sprintf("/tmp/sk_%s.csv", string(buf))
}

func SetHTTPProxy(proxy string) {
    if proxy != BLANK_LINE {
        os.Setenv(HTTP_PROXY, proxy)
    }
}

func Catch() {
    if r := recover(); r != nil {
        fmt.Printf("--\nAn error encountered\n%v\n", r)
    }
}

func Dump(values []string) {
    for i, v := range values {
        fmt.Printf(" [%d]\t%s\n", i, v)
    }
    fmt.Println(" [#]\tend of dump")
}
