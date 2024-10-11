package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	// a[start:stop]

	// This will return an array of elements at index 1 and 2 and will exclude the index 3
	b := a[1:3]                  // Slicing an array teturns slice not an array
	fmt.Printf("%v, %T\n", b, b) // [2,3] []int

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // same as s1[2:len(s1)] [3,4,5,6], when we don't have a stop index means to the full length of an array
	fmt.Println(s1[:3]) // same as s1[0:3]
	fmt.Println(s1[:])  // same as s1[0:len(s1)]
	// fmt.Println(s1[:100]) //Error

	// This will add 100 after the index 4
	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	// This will add 200 at the index 4 and will replace the previous value at the index 4 with 200
	s1 = append(s1[:4], 200)
	fmt.Println(s1)
}
