NAME := dicon/examples/helloworld
VERSION := $(shell git tag -l | tail -1)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'
PACKAGENAME := github.com/recruit-tech/dicon/examples/helloworld

.PHONY: setup main all

all: main

main:
	go build -ldflags "$(LDFLAGS)" $(PACKAGENAME)/cmd/hello

## Install dependencies
setup:
	go get -u github.com/recruit-tech/dicon/cmd/dicon

## generate DI Container
dicon:
	dicon generate --pkg di,component
