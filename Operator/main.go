// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	/*a, b := 4, 2
	r := (a + b) / (a - b) * 2
	fmt.Println(r)

	r = 9 % a
	fmt.Println(r)

	// assignment operator
	// increment assignment
	a += b

	fmt.Println(a)*/

	a, b := 2, 3
	//decrement assignment
	b -= 2
	fmt.Println(b)

	//multiplication assignment
	b *= 10
	fmt.Println(b)

	//division assignment
	b /= 5
	fmt.Println(b)

	//Modulus assignment
	a %= 3
	fmt.Println(a)

	x := 1
	x += 1
	x++
	fmt.Println(x)

	x--
	fmt.Println(x)
	// fmt.Println(5 + x--) // "++ and -- is not an operator in go". So using "--" will give an error

}
