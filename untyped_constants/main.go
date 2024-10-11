package main

import "fmt"

func main() {
	// typed constants
	const a float64 = 5.1

	// this is an untyped constant
	const b = 6.1

	// Const expressions always initialized at the compile time
	const c float64 = a * b
	const str = "Hello, " + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	// const x int = 5
	// Int cannot be multiplied by a float because go is a strong typed language
	// const y float64 = 2.2 * x

	// untyped constant doesn't have a fixed type and it is not forced to obey the type rules
	const x = 5
	const y = 2.2 * 5 // we can multiply float by int because both x and y are untyped constants
	fmt.Printf("%T \n", y)

	var i int = x     //x changes to int
	var j float64 = x // behind the scenes var j float64 = float64(x)
	var p byte = x

	fmt.Println(i, j, p)

	const r = 5
	var rr = r
	fmt.Printf("%T", rr)

	// const are less restrictive than variables
}
