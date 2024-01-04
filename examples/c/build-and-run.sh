#!/bin/bash

# build example
gcc -o example example.c ../../binaries/macos-arm64/poseidon.so &&

# execute example
DYLD_LIBRARY_PATH=DYLD_LIBRARY_PATH:../../binaries/macos-arm64/ ./example