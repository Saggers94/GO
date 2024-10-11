package main

// import "fmt"

// import "fmt" means error

// import packages are file scoped
// permitted in go
import (
	"fmt"
	f "fmt"
)

const done = false // package scoped
// This will not give an error
var b int = 10

// Package scoped
func main() {
	var task = "running" // local block scoped
	fmt.Println(task, done)

	// this is permitted
	const done = true // local block scoped
	f.Printf("done in main() is %v\n", done)
	f1()

	f.Println("Bye bye!")
}

func f1() {
	const done = true
	fmt.Printf("done in f1(): %v\n", done) // this is done from package scope

	// unused variable can cause issues so use blank identifier with caution, mainly in testing
	a := 10
	_ = a
}
