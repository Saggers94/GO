package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano() / 1000000
	// select is only used with channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut!"
	}()

	// If we have add non-blocking default clause in the select statement
	// than we would need this timeout to make sure that the channels are not empty
	time.Sleep(time.Second * 2)
	for i := 0; i < 2; i++ {
		// Select is just like a switch but it is only used with channels
		// Select will block the main goroutine until one of its channel has data to send
		// we can add select statement for non-blocking channel operation
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		default:
			fmt.Println("No activity")
		}
	}

	end := time.Now().UnixNano() / 1000000

	fmt.Println(end - start)
}
