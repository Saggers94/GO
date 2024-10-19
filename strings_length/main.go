package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1)) //6 bytes

	name := "Codruța" // 8 bytes, this is because the rune 'ț' is encoded with 2 bytes
	fmt.Println(len(name))

	fmt.Println(len("£")) // 2 bytes

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("£")

	fmt.Println(n, m)
}
