package main

import "fmt"

func main() {
	var x = 3   // int type
	var y = 3.1 //float type

	// x = x * y // compile type error
	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5 // int type
	fmt.Printf("%T\n", a)

	// Go allows to have a same variable name with a different type
	var b int64 = 2
	a = int(b)
}
