package main

import "fmt"

func main() {
	// There is no connection with variables title, author and year
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	// much more better approach is to use Struct
	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string //Because title and author are the same type we can write it like this
		year          int
	}

	// Struct literal, here the book is struct type and myBook is a struct
	// it automatically assigns the values to the struct fields,
	// for example, title will be "The Devine Comedy", author will be "Dante Aligheri" and year will be 1320
	// So here order matters because for the right information you are 100% reliable on the order
	// which is not recommended way
	myBook := book{"The Devine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	// Now the below code would work as well
	// bestBook := book{title: "Animal Farm", author: "George Orwell"}
	_ = bestBook

	// the fields that we haven't specified will be set to their zero values
	// empty string for author and zero for year
	aBook := book{title: "Just a Book"}
	fmt.Printf("%#v\n", aBook)
}
