package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	// Checking if the substring is present in the strin
	result := strings.Contains("I love go programming!", "love")
	p(result) // true

	result = strings.Contains("I love go programming!", "lovex")
	p(result) // false

	// Check if x or y is present in the string
	result = strings.ContainsAny("success", "xy")
	p(result) //false

	// Check if x,y or s is in the string
	result = strings.ContainsAny("success", "xys")
	p(result) //true because s is present in the string

	// Single quotes for rune
	result = strings.ContainsRune("golang", 'g')
	p(result) //true

	// How many occurances of a e in the "Cheese"
	n := strings.Count("Cheese", "e")
	p(n)

	// if substring is an empty string than Count returns the length of the string plus one
	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("Go PyThon JaVA"))
	p(strings.ToUpper("Go PyThon JaVA"))

	p("go" == "go") //true
	p("Go" == "go") //false

	// ToLower functions loops through each rune and than converts it to a lower case
	// which is not efficient. So it is better to use the EqualFold function
	p(strings.ToLower("Go") == strings.ToLower("go")) //true

	// Recommended way to compare strings
	p(strings.EqualFold("Go", "go")) //true

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	// This will replace first two dots with colons
	// myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // with "-1" it will replace all the dots
	p(myStr)

	// ReplaceAll will replace all occurances of the dot by colon
	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	// Split will give back the slice of the type, in this case it would be slice of the string
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	// When we provide the empty string as the character than it will split the string
	// by every rune literal
	s = strings.Split("Go for go!", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	// myStr = strings.Join(s, "-") // joining the string with the hyphen
	myStr = strings.Join(s, "XXXX")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // this will sploit the string by white spaces and new lines
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	// TrimSpace will remove all leading tabs and white spaces
	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux! \n")
	fmt.Printf("%q\n", s1)

	// Trim will remove the given characters from the string
	// in this case, it is all dots and exclamations marks
	s2 := strings.Trim("... Hello, Gophers!!!!!", ".!")
	fmt.Printf("%q\n", s2)
}
