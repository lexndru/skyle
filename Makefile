GOBUILD=go build
GOCLEAN=go clean
GOGET=go get
GOTEST=go test
BIN_NAME=skyle
VERSION=0.0.9alpha
BUILD=`git log -1 --format="%H"`
OSARCH=`uname`/`uname -i`
FLAGS=-ldflags "-X main.SKYLE_VERSION=$(VERSION) -X main.SKYLE_BUILD=$(BUILD) -X main.SKYLE_OSARCH=$(OSARCH)"

.PHONY: build all clean deps install test

all: clean build install test

build:
	$(GOBUILD) $(FLAGS) -o $(BIN_NAME) -v -x

clean:
	$(GOCLEAN) -x -v
	rm -f $(BIN_NAME)

deps:
	$(GOGET) -v -x

install:
	mv $(BIN_NAME) /usr/bin/$(BIN_NAME)

test:
	$(GOTEST) -v -x -cover
