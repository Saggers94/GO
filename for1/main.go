package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// We don't have a while loop in go but we can use for loop to act as a while loop
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// The below loop is an infinite loop
	sum := 0
	for {
		sum++
	}
	fmt.Println(sum) // this line is never reached
}
