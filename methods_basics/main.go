package main

import (
	"fmt"
	"time"
)

type names []string

// Any value of type names does get access to this method
// In the definition of a method extra parameter is added that has the named type
// and than the function name is added
// Here "n" is called method receiver
// "n" is like this or self from an object oriented programming
// print is called method or receiver in GO
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) //time.Duration

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds) //float64
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Dan", "Maria"}
	friends.print() //Here friends have the type names and so it can call the method print on it with a dot
	// This is like calling this.print() in an object oriented programming
	// since we do not have a class in Go, we use named types to achieve the same thing

	//Another way to call a method is
	names.print(friends)

	// The main difference between method and function is
	// method belongs to a type and function belongs to a package

	// we can convert one type to another only if they share same underlying type
	var n int64 = 93213913
	fmt.Println(n)
	fmt.Println(time.Duration(n)) // convert n to time.Duration // Here Duration is a method
}
