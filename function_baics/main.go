package main

import (
	"fmt"
	"math"
)

// go functions parameter doesn't pass by reference
// go functions parameter pass by value always
func f1() {
	fmt.Println("This is f1() function")
}

// Behind the scenes the function create two int variables and assign the values to them
func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// This is called shorthand parameter
// where a,b and c are int type
// d and e are float64 type
// s is string type
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// This function will return a float64 value
func f4(a float64) float64 {
	return math.Pow(a, a)
}

// This function will return two int values
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// having multiple named return type is permitted in go
// func sum(a, b int) (s, a int, z float64, s, y string)

// named return value with "s"
func sum(a, b int) (s int) {
	fmt.Println(s) // s = 0 at this point in time which is a zero value of int

	s = a + b
	// we can return the named parameter without any value, it will return s automatically
	// This is known as "naked return"
	// Naked return should only be used in small functions they harm the readability in the longer function
	return
}

func main() {
	// Go automatically executes only main and init functions which are predefined function names
	// f1() // nothing will happen because f1() is not called in main function

	f1() //functions body will be executed
	f2(5, 7)
	f3(1, 2, 3, 4.4, 5.5, "Go")
	p := f4(5.1)
	fmt.Println(p)

	// If i am not interested in the second value than I have to use blank identifier "_"
	// otherwise the compiler will throw an error
	// a, _ := f5(8, 9)

	a, b := f5(8, 9)
	fmt.Println(a, b)

	mySum := sum(5, 6)
	fmt.Println(mySum)
}
