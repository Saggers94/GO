package main

import "fmt"

func main() {
	var employees map[string]string

	// Zero value of a map is nil
	fmt.Printf("%#v", employees)

	fmt.Printf("No of Pairs:%d\n", len(employees))

	// "Dan" key is not present in the map it returns the zero value for that type
	// Empty string for string, zero for int, false for bool
	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	// Zero value for a float will be returned
	fmt.Printf("%#v\n", accounts["0x323"])

	// It is not possible to use slices as keys in a map as it is not comparable
	// we can only use comparable types as keys in a map
	// var m1 map[[]int]string

	// Arrays are comparable and can be used as keys in a map
	// Comparable meaning two elements of the same type can be compared using == operator
	var m1 map[[5]int]string
	_ = m1

	// This will give run time error as the map is nil
	// We have to initialize a map before adding elements to it
	// employees["dan"] = "Programmer"

	// with "{}" in the end we are initializing the map, so it is not nil anymore
	// it is an empty map and so we can add elements to it
	people := map[string]float64{}
	people["john"] = 21.4
	people["Marry"] = 22.5
	fmt.Println(people)

	// with make function you can create a map that is initialized but an empty map
	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 233.4,
		"EUR": 555.4,
		// 50:66.6, error all keys must be the same type
		"CHF": 323.4, // last comma is mandatory when we are adding elements on multiple lines
	}

	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println(m)

	// When we use the same key from the map than it will update the value
	balances["USD"] = 500.0
	// When we don't use the same key than it will add a new key value pair
	balances["GBP"] = 100
	fmt.Println(balances)

	//this will give 0 because the key is not present in the map and so
	// it will return the zero value of the value type
	fmt.Println(balances["RON"])

	balances["RON"] = 0
	// When we use "v,ok" than it will return the value and a boolean
	// if the key is present than the "ok" will be true otherwise false
	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON balance is:", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map ")
	}

	// Maps are designed in go for a quick look up and not for looping
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	// This will delete the key value pair of "USD" from the map
	delete(balances, "USD")
	fmt.Println(balances)
}
