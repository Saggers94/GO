package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ //overflow, x is 0 which is a minimum value of uint8
	fmt.Println(x)

	// a := int8(255 + 1) //if the expression executes at a compile time than go catches the overflow

	// when int8 overflows its maximum value, it goes to the minimum value of its type
	var b int8 = 127
	fmt.Printf("%d\n", b+1) // this overflow error will not be caught at the compile time because the expression does not run at a compile time

	// When int8 overflows its minimum value, it goes to the maximum value of its type
	b = -128
	b--
	fmt.Println(b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)

	// Float number overflows to infinity
	f = f * 1.2
	fmt.Println(f)

	// Constant does not overflow because we cannot change it
	// const i int8 = 300 , it gives a compile time error

}
