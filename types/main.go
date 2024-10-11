package main

import "fmt"

func main() {

	// INT type
	var i1 int8 = -128

	// var i2 int8 = -129 // this will give an error cuz constant -129 overflows int8

	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	// var i2 uint16 = 65539 // this will give an error
	fmt.Printf("%T\n", i2)

	// FLOAT type
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// Additional integer types called byte and rune
	// byte is an alias for uint8
	// rune is an alias for int32

	// Rune type
	var r rune = 'a'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	// Bool type
	var b bool = true
	fmt.Printf("%T\n", b)

	// String type
	var s string = "Hello, Go!"
	fmt.Printf("%T\n", s)

	// Array Type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// Slice type
	// slice is like an array but it can shrink or grow
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// Map type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}

	fmt.Printf("%T\n", balances)

	// Struct type
	// is almost same as class concept in object oriented programming

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// Pointer type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// Function type
	fmt.Printf("%T\n", f)
}

func f() {}
