package main

import (
	"fmt"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
	"strconv"
)

func main() {
	for i := 1; i < 257; i++ {
		if i%20 == 1 {
			fmt.Println()
		}
		fmt.Print(style.New().Bg(color.New(i - 1)).Render(pad(i - 1)))
	}
	fmt.Println()

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

func pad(i int) string {
	if i < 10 {
		return strconv.Itoa(i) + "  "
	}
	if i < 100 {
		return strconv.Itoa(i) + " "
	}

	return strconv.Itoa(i)
}
