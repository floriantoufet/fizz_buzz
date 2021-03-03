#!/usr/bin/env bash
go mod vendor
BIN_PATH="bin/main.go"
env GOOS=linux GOARCH=amd64 go build -o fizz_buzz_linux-amd64 ${BIN_PATH}
env GOOS=windows GOARCH=amd64 go build -o fizz_buzz_windows-amd64.exe ${BIN_PATH}
env GOOS=darwin GOARCH=amd64 go build -o fizz_buzz_darwin-amd64 ${BIN_PATH}
