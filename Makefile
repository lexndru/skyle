GOBUILD=go build
GOCLEAN=go clean
GOGET=go get
BIN_NAME=skyle
VERSION=0.0.9alpha
BUILD=`git log -1 --format="%H"`
OSARCH=`uname`/`uname -i`
FLAGS=-ldflags "-X main.SKYLE_VERSION=$(VERSION) -X main.SKYLE_BUILD=$(BUILD) -X main.SKYLE_OSARCH=$(OSARCH)"

all: build
build:
	$(GOBUILD) $(FLAGS) -o $(BIN_NAME) -v -x
clean:
	$(GOCLEAN) -x
	rm -f $(BIN_NAME)
deps:
	$(GOGET) -v
install:
	mv $(BIN_NAME) /usr/bin/$(BIN_NAME)
.PHONY: build all clean deps
