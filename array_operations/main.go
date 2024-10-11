package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	// This modifies the first element of the array
	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// numbers[5] = 100 // This will not work as the array has only 3 elements

	for i, v := range numbers {
		fmt.Println("Index: ", i, "Value:", v)
	}

	fmt.Println(strings.Repeat("#", 10))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index: ", i, "Value:", numbers[i])
	}

	// two dimensional array with 3 elements in each row
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}

	fmt.Println(balances)

	// Two arrays are equal if both of them have the same values in the same order with same elements
	m := [3]int{1, 2, 3}
	// When we assign an array to another variable, it creates a new array with the same values, but its not true for slices because slices are connected
	n := m
	fmt.Println("n is equal to m:", n == m)

	// Modified the m array
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)
}
