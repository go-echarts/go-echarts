#!/bin/bash
basePath=$(cd `dirname $0`; pwd)
templatesPath=${basePath}/templates
datasetsPath=${basePath}/datasets
cd $GOPATH/src/github.com/gobuffalo/packr/packr
go run main.go -i ${templatesPath}
go run main.go -i ${datasetsPath}