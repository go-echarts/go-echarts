@echo off
set basePath=%CD%
set templatesPath=%basePath%\templates
set datasetsPath=%basePath%\datasets
cd %GOPATH%\src\github.com\gobuffalo\packr\packr
go run main.go -i %templatesPath%
go run main.go -i %datasetsPath%