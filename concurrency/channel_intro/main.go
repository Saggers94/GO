package main

import "fmt"

func main() {
	var ch chan int

	fmt.Println(ch) // nil - zero value of a channel

	ch = make(chan int)
	fmt.Println(ch) // channel stores address, passing channel is same as passing the pointer in the function

	c := make(chan int) // new channel
	// channel is a two way communication object, it has sender and receiver

	// <- cahnnel operator

	// Send 10 to the channel
	c <- 10

	// Receive from the channel
	// This means we receive value from the channel, we wait to the channel have the value and when we have it
	// we assign it to the variable, in this case it's num
	num := <-c

	// channel operator indicate the direction of the data flow
	fmt.Println(<-c)

	_ = num

	close(c)
	// channel is always used in junction with Go routine
}
