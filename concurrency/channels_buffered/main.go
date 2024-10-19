package main

import (
	"fmt"
	"time"
)

func main() {

	// Until the channel is full or empty the main goroutine will not be blocked
	// Also it has the first in first out queue descipline
	// c is a buffered channel because it has the capacity of 3 at the initialization
	c := make(chan int, 3) // buffered channel

	go func(c chan int) {
		for i := 1; i < 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i // this line will not block the main goroutine until it has an empty slot
			//in it, after it is full the main will wakes up and receives the data
			fmt.Println("func goroutine #%d after sending data into the channel\n", i)
		}
		// it is not important to close the channel but when necessary
		// here it is necessary because we want to tell goroutine than all the data has been received
		close(c)
	}(c) // buffered channel

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	// v is the value received from the channel, v := <-c
	for v := range c {
		fmt.Println("Main goroutine received value from the channel:", v)
	}

	// unbuffered channel is a blocking channel which will main goroutine to synchronize with other
	// goroutine that's why unbuffered channel is sometimes called as synchronous channel

	// this will be 0 because the channel is closed but it was an int channel so it will be the zero
	// value of the int
	fmt.Println(<-c)

	//sending a value in the closed channel will make it panic
	c <- 10

	// when to use buffered and unbuffered channel
	//unbufferd gives a stronger synchronization as it blocks the main goroutine until the data is received
	// buffered channel operations are decoupled, when we know the upperbound of the data we are going to
	// send it to the channel, we can use buffered channel of that size and perform all the send operation
	// before the first value received
	// Channel buffering also affect the program performance
}
