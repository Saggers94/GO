package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	c <- f
}

func main() {
	ch := make(chan int)
	// Typically considers the best practice to close the channel
	defer close(ch)

	go factorial(5, ch)

	// This is a blocking operation
	f := <-ch
	// After the main goroutine receives the value from the channel, it wakes up prints the message and exit
	// the main goroutine
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 5; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			c <- f
		}(i, ch)

		// Whenever channel send the value it will make the main goroutine to wake up
		// and only after the channel sends the value the main goroutine will print the message
		// and exit, before this we were using waitGroups to wait for all the goroutines to finish
		// before main and also we were using mutex and the lock and unlock method of mutex
		// to solve the data race issue
		fmt.Printf("Factorial of %d is: %d\n", i, <-ch)
	}

}
