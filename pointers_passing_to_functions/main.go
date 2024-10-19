package main

import "fmt"

// typing of a pointer is done by adding * in front of the type
func change(a *int) *float64 {
	// Functions works on copies not on the originals in Go
	// And that's why it is pass by value and not pass by reference
	*a = 100 // It changes what a points to, I am not modifying a

	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 66
}

func main() {
	x := 8
	p := &x

	fmt.Println("value of x before calling change():", x) // 8
	change(p)
	fmt.Println("value of x after calling change():", x) // 100

	// function would not change the value of x if we are passing
	// int, bool, float or struct in the argument
	// because the function would have worked on copies not on the original
	// However things work differently for slices and maps
	// So in order to modify the variable from the function we would need to pass the pointer
	// that changes the value on that memory address
	fmt.Println("value of x before calling changeVar():", x) // 100
	changeVar(x)
	fmt.Println("value of x after calling changeVar():", x) // 100
}
