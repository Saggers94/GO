package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python")
	// There is an invisible break statement here so, we don't need to add a break statement after evey case
	// break

	case "Go", "golang":
		fmt.Println("Good, go for Go! you are using curly braces {}!")

	default:
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	switch true {
	// The first case will give an error as I cannot compare a bool with an int
	// case n:
	// 	fmt.Println("This will not print")
	case n%2 == 0:
		fmt.Println("Even number")

	case n%2 != 0:
		fmt.Println("Odd number")

	default:
		fmt.Println("Never here")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)
	// Missing expression after switch keyword means "true"
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
