#!/bin/sh

STAGED_GO_FILES=$(git diff --cached --name-only -- '*.go')

if [[ $STAGED_GO_FILES = "" ]]; then
    echo "no updated go files"
else
    for file in $STAGED_GO_FILES; do
        go fmt $file
        git add $file
    done
fi

golang-ci run