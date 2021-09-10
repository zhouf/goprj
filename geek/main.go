package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	println("Hello Go!")
	var s = "好好学习"
	println(len(s))
	println(utf8.RuneCountInString(s))

	fmt.Println("hhhhhh")
}
