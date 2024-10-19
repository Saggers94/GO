package main

import (
	"fmt"
	// import "fmt" // this will also give an error
	f "fmt"
)

// so we have to use it with an alias

// There are three scopes in GO, Package, File and local(or blocked)

func main() {
	fmt.Println("Scope means name visibility")

	// I can call the function directly without using the package. Because it is in the same package. Even
	// though it is in a different file
	sayHello("Gopher")
	f.Println("Hello from f")
	// With go run *.go, we can run all the files in the package

	// const, var and const are package scoped, so, available to all the files
	// but the variable that is defined in the function is only available to that function
	tf := toFahrenheit(boilingPoint)
	fmt.Println("Boiling point of water in Degrees Fahrenheit:", tf)
}

// This will give an error as we can not define the same function in the package scope
// func sayHello(name string) {
// 	fmt.Printf("Hello %s!\n", name)
// }
