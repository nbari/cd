.PHONY: all get test clean build cover compile goxc bintray

VERSION=0.0.1
GO ?= go
BIN_NAME=cd
GITHASH=$(shell git rev-parse HEAD)

all: clean build

get:
	${GO} get

build: get
	${GO} build -ldflags "-X main.version=${VERSION} -X main.githash=${GITHASH}" -o ${BIN_NAME}

clean:
	@rm -rf ${BIN_NAME} ${BIN_NAME}.debug *.out build debian

test: get
	${GO} test -v
