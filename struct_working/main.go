package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	// If we try to access the field that struct doesn't have than we will get a compile error
	// page := lastBook.pages

	// Struct in go are fixed at the compile time, you cannot add or remove field on the run time
	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878

	// Using the + modifier to print both the field name and field value
	fmt.Printf("%+v\n", lastBook)

	// The map can't be compared to another map only to nil just like a slice
	// we can compare struct values using == operator

	// aBook := book{title: "Anna Karenina", author: "", year: 0}
	// fmt.Println(aBook == lastBook) // false because author and year are the zero values

	aBook := book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}
	fmt.Println(aBook == lastBook) // true because all the field values are the same

	// using the = sign, Slices and map doesn't create a copy of the data but a reference to the data
	//creating a copy of the struct

	// with the equal sign we are able to create a copy of the struct unlike slice and maps
	// so if we change the field value of the myBook, it will not change the field value of the aBook
	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook) // only the myBook has been modified the other one still contains the same value

	// Anonymous struct : without defining struct type alias
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Spencer",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's Age: %d\n", diana.age)

	// Anonymous fields
	type Book struct {
		string
		float64
		bool
	}

	// We didn't specify the field names but go created the fields based on the type of the field
	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	// We can access the fields by their type
	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool   // anonymous field
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)

	e.bool = true
	fmt.Printf("%#v\n", e)
}
