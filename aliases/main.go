package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a
	_ = b

	// This is an alias because we are using the second type name as an alias to uint
	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d \n", hour/60)
}
