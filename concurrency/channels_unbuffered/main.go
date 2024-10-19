package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbuffered channel

	c2 := make(chan int, 3) // buffered channel
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10 // this line will block the main goroutine until the main wakes up and receives the data
		fmt.Println("func goroutine after sending data into the channel")
	}(c1) //sending in unbuffered channel

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine starts receiving data")
	d := <-c1
	fmt.Println("Main goroutine received data:", d)

	time.Sleep(time.Second)
}
