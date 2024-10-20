package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The Length of a5 is %d\n", len(a5))

	// We use the ellipsis operator when we don't know how many elements are going to be there in the array
	a6 := [...]int{
		1,
		2,
		3,
		4,
		5, //when we have the array values in a different line, It is mandatory to put a comma after the last element
	}
	fmt.Printf("%#v\n", a6)
}
