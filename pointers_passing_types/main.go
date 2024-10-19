package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	// The below assignments will not change the original values cuz they are passed by value
	// which means the function works on the copies not on the original values
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

// * before the type means pointer declaration of that type, *int means a pointer to an int type
// *float means a pointer to a float type
// *string means a pointer to a string type
// *bool means a pointer to a bool type
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	// * before the variable means the value at the address stored in the pointer variable
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300, shortcut offered by Go
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValues():", quantity, price, name, sold)

	// this will change the original values because we are passing pointers
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer():", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift) // It won't change the original values but than how to modify the struct
	fmt.Println(gift)   // {100 Watch}

	// To modify the struct we use pointers and pass the pointer
	fmt.Println("Before calling changeProductByPointer():", gift)
	changeProductByPointer(&gift)
	fmt.Println("After calling changeProductByPointer():", gift)

	// So, int, float, string, bool and struct can be modified only if we pass the pointer

	//However things work differently for slices and maps
	// because slices and maps don't store the actual data but the reference to another memory address
	// where the data is stored which means even though they are passed by value,
	// the copies would point to the same data as the originals
	// when function changes slice or a map, the actual data will change

	prices := []int{1, 2, 3}
	// Slice already contains a pointer to the backing array so we don't need to use pointer to
	// modify the slice inside the function
	changeSlice(prices)
	fmt.Println("Prices slice after calling the ChangeSlice", prices) // [2,3,4]

	// the same is true for maps as well
	// Slices and maps are not to be used with pointers
	myMap := map[string]int{
		"a": 5,
		"b": 12,
	}
	changeMap(myMap)
	fmt.Println("myMap after calling the changeMap:", myMap) // map[a:10 b:20 c:30]

	// array also behaves the same but it is not a good practice to pass the array as an argument
	// to a function, use slices instead

	// passing by value is cheaper than passing by ponters
	// because pointers put an additional pressure on the garbage collector
	// pass pointers to a functions only when it is necessary

}
