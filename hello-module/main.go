package main

// To add this import to a go.mod file which is just like a package.json file in node
// we can use go get github.com/ddadumitrescu/hellomod
import "github.com/ddadumitrescu/hellomod"

func main() {
	hellomod.Salut()
	hellomod.SayHello()
}

// Also go has created a file called go.sum to save the checksum of downloaded dependencies
// when you share your module go.sum plays an important role by validating the checksum of each module
// go.sum has "Module-Path + Version + Hash of Checksum"

// Where are go modules stored?
// Go modules are stored in the $GOPATH/pkg/mod/ directory. This is where go is saving module cache.
