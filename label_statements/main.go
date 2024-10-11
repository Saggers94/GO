package main

import "fmt"

func main() {
	// even though we have an outer variable here, the outer label does not conflict with the outer variable
	// so label doesn't conflict with the variable name
	outer := 19
	_ = outer
	people := [5]string{"John", "Paul", "George", "Ringo", "Pete"}
	friends := [2]string{"Mark", "John"}

	// The "outer" label lives in another space. It is a specific point of the code that we are keeping an eye on
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				// With a lable "outer" after the break we will jump out of the loop of outer keyword
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}
