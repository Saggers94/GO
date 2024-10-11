package main

import "fmt"

func main() {
	count := 0

	// Here "true" will make the for loop run forever, infinite loop
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 {
			// break keyword will break out of the loop and the remaining program will continue to execute
			break
		}
	}

	fmt.Println("Just a message after the loop")
}
