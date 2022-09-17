#!/usr/bin/env just --justfile
# https://vuepress-theme-hope.github.io/v2/zh/faq/vite.html
# https://github.com/markdown-it/markdown-it-emoji/blob/master/lib/data/full.json
# https://codybontecou.com/tailwindcss-with-vitepress.html

run:
    vitepress dev docs
    
build:
    vitepress build docs

go lib:
    go run {{lib}}/main.go

#recode term
rec:
    rm -rf demo.cast
    asciinema rec demo.cast


cast row="15":
    asciicast2gif -h {{row}} demo.cast demo.gif