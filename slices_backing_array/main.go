package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}

	s3, s4 := s1[0:2], s1[1:3]
	// Even though I modified just s3 here, s1 and s4 will also be modified
	s3[1] = 600

	// the Slicing expression doesn't create a new Backing array for a slice
	// So, if we change an element of a variable that has the same backing array than it will change original
	// backing array
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}

	slice1, slice2 := arr1[0:2], arr1[1:3]

	// When we modified the array we are modifying the backing array so all of the slices
	// that we have created through slice expressions will change as well cuz all of them
	// are using the same backing array. For that they are using Pointer
	arr1[1] = 2
	fmt.Println(slice1)
	fmt.Println(slice2)

	// To create a complete new slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// When we use append function, it creates a slice with a new backing array
	// We also have to use elipsis operator to append all of the elements of the slice
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	n1 := []int{10, 20, 30, 40, 50}
	newSlice := n1[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // 3 5
	// Slicing creates a new Slice header but the backing array is the same
	newSlice = n1[2:5]

	// here the capacity would be three because we have done the slicing operation above
	// where we have three elements and since the next slice would have the capacity as address header it would be three
	fmt.Println(len(newSlice), cap(newSlice)) // 3 3 ,

	// This will print the slice header address when we use "&"
	fmt.Printf("%p\n", &n1)
	// Once we slice the n1, "newSlice" will have the different header address
	fmt.Printf("%p %p\n", &n1, &newSlice)

	// Even though we just changed the slice's 0th index element
	// because newSlice and n1 share a same backing array, It will change the n1 as well
	newSlice[0] = 1000
	fmt.Println("n1: ", n1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	// Slice occupies the less memory than an array
	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a)) // size : 40
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(s)) // size : 24

}
