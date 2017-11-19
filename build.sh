#!/bin/sh

# compile AMD64 for Linux, OSX, and Windows
env GOOS=darwin GOARCH=amd64 go build  -o /tmp/mock-email-list-darwin-amd64  .
env GOOS=linux GOARCH=amd64 go build   -o /tmp/mock-email-list-linux-amd64   .
env GOOS=windows GOARCH=amd64 go build -o /tmp/mock-email-list-win-amd64.exe .
