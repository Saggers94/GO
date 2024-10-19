package main

import "fmt"

func main() {
	name := "Sagar"
	// & is used for pointer, pointer is a variable that stores the memory address of another variable
	fmt.Println(&name) // This will give an address in memory where the name variable is stored

	var x int = 2
	ptr := &x // address of x, the type of ptr is *int

	fmt.Printf("ptr is of type %T, and value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	// "*float64" is a type of a Pointer variable of a type float64
	var ptr1 *float64 // zero value is nil
	_ = ptr1

	// another way to create a pointer
	p := new(int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T, and value of %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x) // %p is used to print the address of the variable

	// * is called dereferencing operator
	// * means take the memory address and give me the value of that memory address
	// *p is used to get the value of the variables whoes address is stored in the pointer p
	*p = 90            // equivalent to x = 90
	fmt.Println(x, *p) // x and *p is the same thing
	fmt.Println("*p == x:", *p == x)

	// *float64 and *p have totally different meaning
	// *p means a value whose address is stored in the variable p
	// *float64 means type description, we are declaring a pointer to that type

	*p = 10     //  equivalent to x = 10
	*p = *p / 2 // x divided by 2
	fmt.Println(x)

	// if we have a value stored in x
	// than &x will give the address of x which is a pointer
	// &value => pointer
	// *pointer => value
}
