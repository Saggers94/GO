package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	// Append, add an element at the end of the slice and than return a new slice
	// It doesn't append and return the same slice
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	// We can append many numbers at once and save the new slice in the same variable
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	n := []int{100, 200}
	// With the "n..." we are appending all of the elements of the slice n to a numbers slice
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	// dst := make([]int, len(src))
	dst := make([]int, 2)
	// Copy function copies the elements from the source slice to the destination slice
	// "nn" is the number of elements copied
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
