package main

import "fmt"

func main() {
	var age = 30
	fmt.Println("Age: ", age)

	var name = "John Doe"
	fmt.Println("Name: ", name)

	// Blank identifier
	_ = name

	// Short declaration ":=" operator only works for block scope
	s := "Learning Go!"
	fmt.Println(s)

	// I cannot use ":=" for already defined vaiables
	// name := "Jane Doe" // This will give an error

	// This will work
	name = "Jane Due"

	// multiple variable declaration
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// When we declare multiple variable if one of them is new than it is okay and it won't give
	// Compilation error
	car, year := "BMW", 2000
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	// _, _ = i, j
	// Swapping the variable values
	j, i = i, j

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}

// We cannot use ":=" outside of the function, it will give a compilation error if it is package scoped
// Go infer variable type so we don't need to specify the type of the variable
