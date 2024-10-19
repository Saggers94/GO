package main

import "fmt"

// Any go type can satisfy an empty interface so, it can represent any value
type emptyInterface interface{}

type person struct {
	info emptyInterface
}

func main() {
	var empty emptyInterface
	empty = 5
	fmt.Println(empty)

	empty = "go"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)

	// Interfaces stores two values dynamic type and dynamic value
	fmt.Println(len(empty.([]int))) // we have to use type assertion to get the value

	// Empty interface is used for the unknown type

	you := person{}
	you.info = "Hello"
	you.info = 20
	you.info = []float64{1.2, 2.3, 3.2}
	fmt.Println(you.info)

	// Empty interface can make the program hard to maintain

	// Empty interfaces check the type on the run time so even though it is easy to use
	// it does come with an issue of type safety and you won't know the error until you run the program
}
