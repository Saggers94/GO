package main

import "fmt"

func main() {
	// pointer can pointer to another pointer as well
	a := 5.5
	p1 := &a
	pp1 := &p1

	fmt.Printf("value of p1: %v, address of p1: %p\n", p1, &p1)
	fmt.Printf("value of p1: %v, address of p1: %p\n", pp1, &pp1)

	// dereferencing the pointer to get the value
	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)

	fmt.Printf("**pp1 is %v\n", **pp1)

	**pp1++ // a will be incremented by 1
	fmt.Printf("a is %v\n", a)

	//Comparing pointers
	// Zero value of pointer of any type is nil
	var p2 *int
	fmt.Printf("%#v\n", p2) //(*int)(nil)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // It will return false because both pointers are pointing to a same
	//value but to a different memory location that stores the different value

	p4 := &y
	fmt.Println(p2 == p4) // This will return true because p2 is also pointing to the address of y
	// and p4 is also pointing to the address of y

	
}
