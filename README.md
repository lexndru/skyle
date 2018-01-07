# Skyle
[![Build Status](https://travis-ci.org/lexndru/skyle.svg?branch=master)](https://travis-ci.org/lexndru/skyle)

Notice: the current version v0.0.9alpha is in an early stage. Bugs and errors might occur.

## Introduction
Skyle is an ambitious open-source scraper with its own declarative language. It's purpose is to create a portable and elegant way to scrape various file formats. Get a grasp about the basics from this tutorial http://skyle.codeissues.net

## Download latest build
```
$ wget https://github.com/lexndru/skyle/releases/download/v0.0.10-alpha/skyle-0.0.10-alpha-linux-x86_64.tar.gz
$ tar -vzxf skyle-0.0.10-alpha-linux-x86_64.tar.gz
$ sudo mv skyle /usr/bin/skyle
$ skyle
     _          _
 ___| | ___   _| | ___
/ __| |/ / | | | |/ _ \
\__ \   <| |_| | |  __/
|___/_|\_\\__, |_|\___|
     |___/

Version: 0.0.10alpha
Build: 3d0efec0e46e65...
OS/Arch: Linux/x86_64
```

## Build from sources

#### System requirements (Debian/Ubuntu)
```
sudo apt-get install build-essential git golang-go libxml2
```

#### Get sources and `make`

```
$ git clone https://github.com/lexndru/skyle.git
$ cd skyle
$ make deps
$ make
$ sudo make install
$ skyle
     _          _
 ___| | ___   _| | ___
/ __| |/ / | | | |/ _ \
\__ \   <| |_| | |  __/
|___/_|\_\\__, |_|\___|
     |___/

Version: 0.0.10alpha
Build: 3d0efec0e46e65...
OS/Arch: Linux/x86_64
...
```

## Usage
Create a file with the following content
```
#!/usr/bin/skyle

# getting started in less than 5 minutes
title tutorial
probe http://skyle.codeissues.net/
output keywords.csv

# extract all Skyle keywords
follow //table//kbd/text()
save keyword

# add keyword usage example
follow //table//kbd/@data-usage
save usage
```

Save it, then `chmod +x yourfile && ./yourfile`

## Documentation

#### title
- Type: (Optional) Header
- Usage: Used to give the profile a name or a short description
- Sample: `title Some human-readable text`

#### probe
- Type: Header
- Usage: Main profile input; provide the file path or URI to desired probe
- Sample: `probe /tmp/index.html`

#### output
- Type: Header
- Usage: Main profile output; provide the file path for output file (default CSV)
- Sample: `output /tmp/results.csv`

#### flags
- Type: (Optional) Header
- Usage: Profile customization settings defined as K[=V] pairs
- Sample: `flags timeit verbose mode=write`

| Flag    | Possible values | Meaning |
| ------- | --------------- | ------- |
| timeit  | [true]/false    | Prints runtime in seconds after execution |
| verbose | [true]/false    | Prints verbose output while running profile |
| mode    | write/[append]  | Append mode adds records on the same output; Write mode creates new file each time |
| mime    | mimetype        | Force profile to handle probe as <mimetype> (default text/html) |
| format  | mimetype        | Force profile to write output as <mimetype> (default text/csv) |
| exec    | [sync]/async    | Set shell exec mode: sync or async (unsupported) |
| proxy   | proxy           | Use given HTTP proxy |
| sync    | integer         | Ensure all instructions return at least <integer> values |
| maxiter | integer         | Change the maximum iteration limit |
| timeout | seconds         | Change HTTP request timeout (seconds) |
| cache   | [true]/false    | Enable or disable cache |

#### agent
- Type: (Optional) Header
- Usage: User-agent for HTTP requests (used if *probe* is an URI)
- Sample: `agent Mozilla/5.0 (compatible; Skyle/1.0; +http://codeissues.net)`

#### node
- Type: Instruction
- Usage: Change the ROOT node (available for XML/HTML *probe*)
- Sample: `node //xpath/@expr`

#### follow
- Type: Instruction
- Usage: Evaluates an XPath expression (available for XML/HTML *probe*)
- Sample: `follow //xpath/@expr`

#### next
- Type: Instruction
- Usage: Loop last declared *node* (available for XML/HTML *probe*)
- Sample: `next node`

#### pattern
- Type: Instruction
- Usage: Evaluates a regular expression and extracts first group
- Sample: `pattern [reg]+expr`

#### replace
- Type: Instruction
- Usage: Replaces recursive regular expression with a literal value
- Sample: `replace [reg]+expr VALUE`

#### remove
- Type: Instruction
- Usage: Removes recursive regular expression
- Sample: `remove [reg]+expr`

#### glue
- Type: Instruction
- Usage: Replaces variables in a given string
- Sample: `glue $name is $age years old`

#### keep
- Type: Instruction
- Usage: Filters against a regular expression
- Sample: `keep [reg]+expr`

#### save
- Type: Instruction
- Usage: Save instruction's result for output
- Sample: `save variable [value]`

#### dump
- Type: Instruction
- Usage: Prints the memory stack at the current calltime
- Sample: `dump foobar`

#### flush
- Type: Instruction
- Usage: Resets memory stack with initial values
- Sample: `flush all`

#### exec
- Type: Instruction
- Usage: Launches a shell command and prints stdout and stderr
- Sample: `exec shell-command-here`

## Known issues:
- Fails to recognize MIME types for local files
- Flush slows down performance for XML/HTML files

## License
Copyright 2017 Alexandru Catrina

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
