package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is Down\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		// this is a blocking call so, main function would wait to receive this from the channel
		// and only than it will exit
		c <- url // sending into the channel
	} else {
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}
}

func main() {
	urls := []string{"https://golang1.org", "https://www.google.com", "https://www.medium.com", "https://www.amazon.com"}

	//1.
	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	fmt.Println("No. of goroutine", runtime.NumGoroutine())

	// Infinite loop that keeps checking if the website is up or down after every 2 second
	// for {
	// 	go checkUrl(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// We can loop over the channel just like the slice
	// this loop will wait until it receives a value from the channel and that it will start to execute it
	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, c)
	// }

	// go is a passbyvalue language
	for url := range c {
		// "u" is a copy of the url, this is what pass by value means
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}

}
