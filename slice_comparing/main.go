package main

import "fmt"

func main() {
	var n []int
	// The below is true because slice has not been initialized
	fmt.Println(n == nil)

	m := []int{}
	// The below is false because slice has been initialized
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 5, 3}
	// Slice can only be compared to nil
	// fmt.Println(a == b)

	// to compare two slices, we have to use for loop and compare each element
	var eq bool = true

	a = nil
	for i, valA := range a {
		if valA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b Slices are equal")
	} else {
		fmt.Println("a and b Slices are not equal")
	}
}
