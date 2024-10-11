package main

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		// goto statement will jump to the label "loop"
		goto loop
	}

	// The below code will give an error because we cannot goto a label that is declared after the goto statement
	// 	goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("somthing here")

	// goto statement are dangerous and should be used with caution
}
