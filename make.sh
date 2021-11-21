#!/bin/bash

export GOARCH="amd64"
export GOOS="linux"

go build -o linux/amd64/


export GOARCH="arm64"
go build -o linux/arm64/


export GOARCH="amd64"
export GOOS="windows"

go build -o windows/amd64/
