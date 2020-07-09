@echo off

set GOOS=windows
set GOARCH=386
rem or set GOARCH=amd64

go build -o kafkacli.exe main.go
