#!/bin/bash

# Define the GOPATH; it's the place where Go will store downloaded modules and built binaries
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"

# build .so
go build -a -buildmode=c-shared -o output/poseidon.so main.go
