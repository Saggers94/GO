package main

import "fmt"

func main() {
	s1 := "I love Golang!"
	fmt.Println(s1[2:5]) // lov

	s2 := "萨加尔尔加" //Sagar in chinese

	fmt.Println(s2[0:2]) // returns two bytes, first at index 0 and second at index 1,
	//returns the unicode representation of those two bytes

	// How to convert the slice of string to slice of rune
	rs := []rune(s2)
	fmt.Printf("%T\n", rs) // []int32 because rune is an alias for int32

	fmt.Println(string(rs[0:3])) // This printed out the runes at index 0,1, and 2

	// Converting the slice of a string to slice of a rune is not efficient, new backing array
	// is created each time
}
