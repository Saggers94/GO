package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil", cities == nil)
	fmt.Printf("cities: %#v\n", cities)
	// nil means it is uninitialized and that is why we can call len function on the slice
	fmt.Println("Length of cities:", len(cities))

	// we cannot assign a value to the slice as it is nil
	// cities[0] = "London"

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	// make is a built-in function to create slices with the proper length
	num := make([]int, 2)
	fmt.Printf("%#v\n", num)

	type names []string
	friends := names{"Dan", "Maria"}

	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is", myFriend)

	friends[0] = "Gabriel"
	fmt.Println("My best friend is", friends[0])

	for index, value := range friends {
		fmt.Printf("Index: %v Value: %v\n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
