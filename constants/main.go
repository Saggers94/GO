package main

import (
	"fmt"
)

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Println("Duration in seconds:", duration*secondsInHour)

	// variable belong to run time while constants belong to compile time
	x, y := 5, 1

	// Even though y is zero, the code will compile because y is a variable and variable executes on run time
	fmt.Println(x / y)

	const a = 5
	const b = 1

	// This will be detected on compile time before the execution so we see the error unlike the variables
	// fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	// min2 and min3 will have the same type and value as min1
	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)

	// constant rules
	// 1. you cannot change a constant
	// const temp int = 100
	// temp = 50

	// 2. you cannot initialize a constant at run time
	// const power = math.Pow(2, 3)

	// 3. you cannot use a variable to initialize a constant
	// because variable belongs to run time and constant cannot be initialized with a runtime value
	// t := 5
	// const tc = t

	// 4. len is a built in function that is available at a compile time unlike math.Pow
	// const l1 = len("Hello")
}
