package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50 // modifying the first element of the slice
}

func sumAndProduct(a ...float64) (float64, float64) {
	// sum := 0     // type is an int here and that's why we are getting an error
	// product := 1 // type is an int here and that's why we are getting an error

	sum := 0.     // use dot after the value to make it float
	product := 1. // use the dot after the value to make it float

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	// Instead of printing the string "Sprintf" will return the string
	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)
	return returnString
}

func main() {
	// variadic funtions take variable number of arguments "..." elipsis makes the function variadic
	// variadic means zero or more arguments
	// if the functions take parameter of different types than only the last parameter can be variadic
	// fmt.Println is a variadic function
	f1(1, 2, 3, 4)
	f1() // when we don't have any argument the value pass to the parameter is nil

	// append is also an variadic function
	nums := []int{1, 2}
	nums = append(nums, 3, 4)
	fmt.Println(nums)

	// You can also pass a slice to a variadic function with the variadic operator
	f1(nums...)

	f2(nums...) // slice will be modified because of the same backing array
	fmt.Println("Nums:", nums)

	s, p := sumAndProduct(2.0, 5., 10., 5.6, 5.6, 87.3)
	fmt.Println(s, p)

	// mix variadic and non-variadic parameters
	info := personInformation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info)

	

}
