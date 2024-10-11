package main

import "fmt"

func main() {
	// Strings are define in double quotes, not in a single quote like javascript and python
	// In the double quotes, string is called string literal
	s1 := "Hi there Go!"
	fmt.Println(s1)

	// We use "\" to escape the double quotes
	fmt.Println("He says: \"Hello\"")

	// Another way to use double quotes inside a string is with backticks
	fmt.Println(`He says : "Hello"`)

	// The backticks are called as raw string because \n would not have any special meaning
	// inside the backticks
	s2 := `I like \n Go!`
	fmt.Println(s2)

	fmt.Println("Price: 100\n Brand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike`)

	fmt.Println(`C:\Users\Dan`)
	fmt.Println("C:\\Users\\Dan") // "\" to escape the backslash

	// Every concatinate creates a new string
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

}
