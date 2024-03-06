#!/usr/bin/env just --justfile

amd64_linux := "GOOS=linux GOARCH=amd64"
amd64_win := "GOOS=windows GOARCH=amd64"


run lib:
    go run {{lib}}/main.go

# Update dependices
update:
  go get -u
  go mod tidy -v
