package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		2: 7,
		0: 5,
	}
	//[5,10,7]
	fmt.Println(grades)

	//[0,0,50]
	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names, len(names))

	// Unkeyed element gets the index from the last keyed element
	cities := [...]string{
		5:        "Paris",
		"London", // index 6
		1:        "NY",
	}
	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
}
