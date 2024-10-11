package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			// Continue statement will skip the rest of the loop and continue with the next iteration
			continue
		}
		// this will not be reached when I is an odd number
		fmt.Println(i)
	}
}
