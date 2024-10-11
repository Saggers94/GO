package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go, World!")
	distance := 60.8 // distance in km
	fmt.Printf("The distance in miles is: %f \n", distance*0.621)
}

// IF you get any error of module than run go mod init

// IF you want to see a lot of details than use "go run -x main.go"
// We usually use the go command to manage GO source code
// When we use "go build" it exports out an executable file with a name of the directory
// To give a specific name to an executable file we use "go build -o <filename>"

// "go run" complies and runs the code
// "go build" compiles the application and produces an executable file

// To create an executable file of the application for the windows "GOOS = windows GOARCH = amd64 go build -o winapp.exe"
// For linux "GOOD = linux GOARCH = amd64 go build -o linuxapp.exe"
// for macos "GOOS = darwin GOARCH = amd64 go build =o macapp.exe"

// Both go install and go build will compile the package in the current directory

// if the package is main, go build will place the resulting executable in the current directory
// and go install will move the executable to GOPATH/bin

// when running go install you use paths relative to GOPATH/src

// "gofmt -w main.go" to format the code
// "goftm -w -l application_structure" to check which files have been modified

// When you are in the go directory you can use "go fmt" to format all the files
