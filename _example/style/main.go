package main

import (
	"fmt"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

func main() {
	fmt.Println(style.New().Foreground(color.Magenta).Render("Magenta"))
	fmt.Println(style.New().Foreground(color.Red).Render("Red"))
	fmt.Println(style.New().Foreground(color.LightBlue).Render("LightBlue"))
	fmt.Println(style.New().Foreground(color.Cyan).Render("Cyan"))
	fmt.Println(style.New().Foreground(color.Aqua).Render("Aqua"))
	fmt.Println(style.New().Foreground(color.HotPink).Render("HotPink"))
	fmt.Println(style.New().Foreground(color.Orange).Render("Orange"))
	fmt.Println(style.New().Foreground(color.FullBlue).Render("FullBlue"))
	fmt.Println(style.New().Foreground(color.Blank).Render("Blank"))
	fmt.Println(style.New().Foreground(color.DarkGray).Render("DarkGray"))
	fmt.Println(style.New().Foreground(color.Gray).Render("Gray"))
	fmt.Println(style.New().Foreground(color.Special).Render("Special"))
	fmt.Println(style.New().Foreground(color.Highlight).Render("Highlight"))
	fmt.Println(style.New().Foreground(color.Subtle).Render("Subtle"))
}
