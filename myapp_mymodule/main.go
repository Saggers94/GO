package main

// Make sure not to use https here
import (
	"fmt"
)

func main() {
	fmt.Println(go_math.geometry.cubeVolume(3))
}

// creating a go mod file to this directory
// go mod init myapp_mymodule

// Getting the geometry package and adding it to the pkg directory
// go get github.com/Saggers94/go_math/geometry
