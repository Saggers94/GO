package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is Down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // http://www.google.com
			file += ".txt"

			fmt.Printf("Writing response body to %s \n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	// Concurrently working on each url with goroutines
	urls := []string{"https://golang1.org", "https://www.google.com/a.html", "https://www.medium.com", "https://www.amazon.com"}

	var wg sync.WaitGroup

	// number of go routines is equal to the number of urls
	wg.Add(len(urls))

	for _, url := range urls {
		// Concurrency means two go routines executing at the same time
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("Number of Goroutines:", runtime.NumGoroutine()) // 5
	wg.Wait()
}
