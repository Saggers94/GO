package main // In go any executable package is called "main" with a main function

import (
	"GO/numbers"
	"fmt"
) // importing a package means importing what's inside the package
// and we import directory not the files
// all the go packages are relative to a path that we have describe in the GOPATH env
// here it is the src package, that's why to import numbers we have to write "mypackages/numbers"

func main() {
	var x int = 40
	fmt.Printf("%d is even: %t \n", x, numbers.Even(x))
	// there could be more than one file in the package but all the files will belong to the same package
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100))
	fmt.Println(numbers.OddAndPrime(25))
	fmt.Println(numbers.OddAndPrime(13))
}
