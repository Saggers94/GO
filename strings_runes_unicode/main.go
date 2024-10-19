package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Rune is represented by single quotes and it is called rune literals
	// And Rune are represented as int32 unicode code points
	// In Go, we don't have a char type so we use rune to represent a character
	var1, var2 := 'a', 'b'

	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	// in the string each character is stored from 1 to 4 bytes
	str := "ţară"         // four characters of a rune in the string but it occupied 6 bytes
	fmt.Println(len(str)) // The length of this will be 6 which is the number of bytes

	// When we use the indexes we get the byte not rune
	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		// "%c" print out the rune or character
		fmt.Printf("%c", str[i]) // This will print out the bytes
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	for i := 0; i < len(str); {
		fmt.Println(str[i:])
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
