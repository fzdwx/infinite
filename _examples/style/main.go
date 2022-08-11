package main

import (
	"fmt"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

func main() {
	fmt.Println(style.New().Fg(color.Magenta).Render("Magenta"))
	fmt.Println(style.New().Fg(color.Red).Render("Red"))
	fmt.Println(style.New().Fg(color.LightBlue).Render("LightBlue"))
	fmt.Println(style.New().Fg(color.Cyan).Render("Cyan"))
	fmt.Println(style.New().Fg(color.Aqua).Render("Aqua"))
	fmt.Println(style.New().Fg(color.HotPink).Render("HotPink"))
	fmt.Println(style.New().Fg(color.Orange).Render("Orange"))
	fmt.Println(style.New().Fg(color.FullBlue).Render("FullBlue"))
	fmt.Println(style.New().Fg(color.Blank).Render("Blank"))
	fmt.Println(style.New().Fg(color.DarkGray).Render("DarkGray"))
	fmt.Println(style.New().Fg(color.Gray).Render("Gray"))
	fmt.Println(style.New().Fg(color.Special).Render("Special"))
	fmt.Println(style.New().Fg(color.Highlight).Render("Highlight"))
	fmt.Println(style.New().Fg(color.Subtle).Render("Subtle"))
	fmt.Println(style.New().BorderStyle(style.DoubleBorder()).Render("hello world"))
	fmt.Println(style.New().Width(200).Center().Render("hello world"))
}
