@echo off

set CGO_ENABLED=0

go build -o kafkacli.exe main.go