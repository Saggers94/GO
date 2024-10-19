package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Using other delimiters with words and also with runes
	// scanner.Split(bufio.ScanWords) // scan the file word by word
	scanner.Split(bufio.ScanRunes) // scan the file rune by rune, character by character
	success := scanner.Scan()

	if success == false {
		err = scanner.Err() //This function will return the error value
		if err == nil {
			log.Println("Scan was completed and it reached the end of the file")
		} else {
			log.Fatal(err)
		}
	}

	//This will return the first line of the file
	// fmt.Println("First line found:", scanner.Text())
	fmt.Println("First line found:", scanner.Bytes())

	// To have the all of the lines of the file
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
