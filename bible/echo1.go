package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("----------------------------------")

	s, sep = "", ""
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	fmt.Println(strings.Join(os.Args[1:], ","))
}
