#!/usr/bin/env just --justfile

run lib:
    go run {{lib}}/main.go

# Update dependices
update:
  go get -u
  go mod tidy -v

#recode term
rec:
    asciinema rec demo.cast

cast:
    asciicast2gif demo.cast demo.gif
    rm -rf demo.cast