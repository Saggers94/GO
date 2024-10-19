package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	// b := map[string]string{"B": "Y"}
	b := map[string]string{"A": "X"}

	// Maps can only be compared to nil
	// fmt.Println(a == b)

	// When the key and value pair are the string type, we can convert the map to string
	// and then compare the strings
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
