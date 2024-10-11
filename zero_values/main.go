package main

import "fmt"

func main() {
	var a = 4
	var b = 5.3

	// a = b this will give an error because a and b are of different types
	a = int(b)

	fmt.Println(a, b)

	/* var x int
	x = "5" // we cannot assign a string literal to an integer type
	*/

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)

	// 	break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var

	// You can edit this code!

	/* fmt.Println("Hello, playground")
	var mv int        // max value
	var max_value int // Not OK

	var packetReceived int // Not OK, name too long
	var b int              // OK

	const MAX_VALUE = 100 // NOt ok
	const N = 100         // Better

	var Max = 100 // if we give this a name of "max", it means that it is not going to be used in other packages
	// but if the name is "Max" than it will be used in the other packages

	// Use camel case in the go

	maxValue := 1000   // recommended
	max_values := 1000 // Not recommended

	writeToDB := true // ok, idiomatic
	writeToDb := true // not ok */

	/* all about fmt package
	name := "Sagar"
	fmt.Println("Hello, playground. I am", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)

	// %d is for int
	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f", 5, 6.8)

	// we cannot have double quotes inside double quotes
	// so, we have to escape the inne double quote with "\"
	fmt.Printf("He says: \"Hello Go!\"")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi
	// \n means a new line
	fmt.Printf("Radius is %d\n", radius)

	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	// %q is for a string
	fmt.Printf("This is a %q", figure)

	// %v -> replaced by any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	// %T -> type
	fmt.Printf("figure is of type %T \n", figure)
	fmt.Printf("radius is of type %T \n", radius)
	fmt.Printf("pi is of type %T \n", pi)

	// %t -> bool
	closed := true
	fmt.Printf("closed is of type %T \n", closed)

	// %b -> base 2
	// 8 bit
	fmt.Printf("%08b \n", 55)
	// 16 bit haxadecimal
	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9

	// only three decimal points with ".3"
	fmt.Printf("x * y = %.3f\n", x*y)
	*/
}
