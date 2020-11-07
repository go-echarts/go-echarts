#!/bin/bash

version="v2.0.2"
app="go-echarts"

function pre_check() {
  info=$(diff -u <(echo -n) <(format))
  if [ -n "$info" ]; then
    echo ">> Make sure you have formatted the codebase before committing."
    echo ">> Use 'git status' command to show what have changed."
    exit 1
  fi
}

function format() {
    go fmt ./... && ls -d */ | xargs goimports -w
}

function test() {
    go test -v ./...
}

function release() {
    pre_check

    echo ">> Building tag..."
    git tag -a $v -m "Release: $v"

    echo ">> Pushing tag to the remote...."
    git push origin $v
}

function help() {
    echo "$0 --cmd [format|test|release]"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "format" ];then
    format
elif [ "$1" == "test" ];then
    test
elif [ "$1" == "release" ];then
    release
else
    help
fi
