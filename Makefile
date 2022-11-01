BIN := bin/wishes-server
HASH := $(shell git rev-parse --short HEAD)
COMMIT_NAME := $(shell git show -s --format=%s ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_NAME})

PKG_ROOT := github.com/xopoww/wishes

.PHONY : generate clean build rebuild test

generate:
	go generate ./...
	go mod tidy

clean:
	go clean

build:
	go build -race -o ${BIN} \
		-ldflags="-X '${PKG_ROOT}/internal/meta.buildVersion=${VERSION}' -X '${PKG_ROOT}/internal/meta.buildDate=${BUILD_DATE}'" \
		./cmd/wishes-server

rebuild: clean generate build

test:
	go test ./...