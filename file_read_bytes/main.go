package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := make([]byte, 2)
	// ReadFull will read the file until the slice is full
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ReadAll reads every thing of a file as a string
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))

	//Entire content of the file in the byte slice
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data)
}
