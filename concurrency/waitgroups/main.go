package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// waitgroup must be passed as a pointer
func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second) // to stimulate a long running process
	}
	fmt.Println("f1 execution stopped")
	wg.Done() // notifying this worker is done
	// (*wg).Done() we can also write this
}

func f2() {
	fmt.Println("f2 execution started")
	for i := 0; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution stopped")
}

// We have one "main" goroutine
func main() {

	// WaitGroup is what we use to have the sync in between the goroutines
	// in this case the "main" goroutine and the f1 goroutine
	// with waitgroup, you can wait for a collection of goroutines to finish
	var wg sync.WaitGroup

	wg.Add(1)

	// define a new goroutine with a go keyword
	go f1(&wg)
	fmt.Println("No of goroutines after f1,", runtime.NumGoroutine())

	f2()

	// As we have added waitGroup we don't need this
	// time.Sleep(time.Second * 2)

	//we use
	wg.Wait() // to wait for the goroutines to finish before the main finishes
	fmt.Println("main execution stopped")
}
