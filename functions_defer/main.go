package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func fooBar() {
	fmt.Println("This is fooBar()")
}

func main() {

	// Why are we using defer?
	// Defer is used to delay the execution of a function until the surrounding function returns
	// It is used for cleanup actions, like closing the file after opening it

	// foo()
	defer foo() // this will defer the execution of this function after the main function is executed
	bar()
	// fooBar()
	fmt.Println("Just a string after deferring foo() and calling bar()")

	// The last function that was deferred will be executed first
	// so in this case fooBar() will be executed first and than foo() will be executed
	defer fooBar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Code that works (read/write) with the file goes here
}
