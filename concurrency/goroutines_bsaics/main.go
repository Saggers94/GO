package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 execution stopped")
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
	fmt.Println("main execution started")
	// Getting the number of CPU cores avaialble to the program
	fmt.Println("No. of cpus", runtime.NumCPU())
	// Number of goroutines that currently exist
	fmt.Println("No of goroutines", runtime.NumGoroutine())

	// Environment variables
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)

	fmt.Println("GoMaxPROCS", runtime.GOMAXPROCS(0))

	// main function doesn't wait for the goroutines to execute
	// or to finish
	// main function terminates before the goroutines executes
	// we want a way to wait for the goroutines to finish before the execution of the main function stops
	go f1()
	fmt.Println("No of goroutines after f1,", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped")
}
