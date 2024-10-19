package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	// idiomatic way is to use defer.
	// defer is used to delay the execution of the function until the surrounding function returns
	// Common way to close a file
	defer file.Close()

	// This is how we convert a string to a byteSlice
	byteSlice := []byte("I learn Golang!") // one ASCII character is one byte
	// so, 15 characters will be 15 bytes
	// Using the Write method of the file, It returns the number of byteWritten and the err value
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Byte written to file: %d\n", byteWritten)

	// quick way to dump a byte slice
	bs := []byte("Go Programming is cool!")
	// ioutil is deprecated
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Adding ioutil or bufio adds the more binary packages on to the executable files
	// So, It is better to use just the os package. Since, we have already use it to open
	// the file
}
