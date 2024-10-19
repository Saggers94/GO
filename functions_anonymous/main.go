package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	// Anonymous function doesn't have any name and that's why it's called right away
	// because we can't call it by any name, It is also called closure
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function")

	// Anonymous function can be used when we want to create a function inside a function
	// In go, we can return a function
	// Also, we can pass a function as an argument to another function

	// A first class function means it allow functions to be assigned to variables, passed
	// as an argument to other functions or return from other functions
	a := increment(10)
	fmt.Printf("%T\n", a) // Type will be "func() int"

	a()              // 11
	fmt.Println(a()) // 12
	fmt.Println(a()) // 13
	fmt.Println(a()) // 14
}
