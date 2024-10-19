package main

import "fmt"

// init function is used to initialize the global variables
//declare an array
var numbers [10]int

// main and init function is a reserved keywords
// init function does not take any arguments and does not return anything and it is called automatically
func init() {
	fmt.Println("This is first init")
}

// You can also define another init
func init() {
	fmt.Println("This is second init")

	// without the init function, you can not initialize the array before main is called
	// and you cannot have a for loop outside the main function
	// so, init gives us an opportunity to initialize the array before main is called
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func main() {
	fmt.Println("This is main")
	// init() you can not call this manually
	// If you want to initialze the array before main is called, you can use the init function
	fmt.Printf("%#v", numbers)
}
