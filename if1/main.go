package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive")
	}

	_ = inStock

	// In go there's nothing like a truthiness of a value like in other programming languages
	// We can use only boolean expression in the if statements
	if price < 100 && inStock {
		fmt.Println("Buy it!")
	}

	// The below code will give a compile time error
	// if price {
	// 	fmt.Println("something")
	// }

	if price < 100 {
		fmt.Println("It's cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's expensive")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("You can not vote!please return in %d years \n", 18-age)
	} else if age == 18 {
		fmt.Printf("This is your first vote! \n")
	} else if age > 18 && age < 100 {
		fmt.Printf("Please vote, it's important")
	} else {
		fmt.Printf("Invalid age")
	}

}
