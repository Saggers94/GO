package main

import (
	"fmt"
	"math"
)

// Interface in go implements automatically as long as we have the methods that are defined in the named type
type shape interface {
	area() float64
	perimeter() float64
}

// circle and rectangle are named types
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// circle and rectangle are implementing the shape interface because we have created methods
// like area and perimeter on the circle and rectangle named types through receiver
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// If I will comment this out than the print function will not work because the rectangle is not implementing
// area method and so it is not implementing the shape interface
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// in printCircle and printRectangle we are repeating ourselves which is not a good practice
// func printCircle(c circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter:", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter:", r.perimeter())
// }

// Because rectangle and circle are implementing the method area and perimeter, they are implementing the shape interface
// And so, we can pass them to the print function
// This function takes in the shape type which means it can take any named type that implements the shape interface
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	fmt.Println("Hello, World!")
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3, height: 2.1}

	// Even though print takes the argument of type shape, we can pass the circle and rectangle to it
	// because circle and rectangle are implementing the shape interface
	// so we pass concrete circle and rectangle values
	print(c1)
	print(r1)
}
