GOBUILD=go build
GOCLEAN=go clean
GOGET=go get
GOTEST=go test
BIN_NAME=skyle
VERSION=0.0.10-alpha
BUILD=`git log -1 --format="%H"`
OS=`uname`
ARCH=`uname -i`
FLAGS=-ldflags "-X main.SKYLE_VERSION=$(VERSION) -X main.SKYLE_BUILD=$(BUILD) -X main.SKYLE_OSARCH=$(OS)/$(ARCH)"
RELEASE=$(shell echo $(BIN_NAME)-$(VERSION)-$(OS)_$(ARCH).tar.gz | tr A-Z a-z)

.PHONY: build all clean deps install test release

all: clean build test

build:
	$(GOBUILD) $(FLAGS) -o $(BIN_NAME) -v -x

clean:
	$(GOCLEAN) -x -v
	rm -f $(BIN_NAME) 2> /dev/null
	rm -f $(RELEASE) 2> /dev/null

deps:
	$(GOGET) -v -x

install:
	mv $(BIN_NAME) /usr/bin/$(BIN_NAME)

test:
	$(GOTEST) -v -x -cover

release:
	rm -f $(RELEASE) 2> /dev/null
	tar -vzcf $(RELEASE) $(BIN_NAME)
