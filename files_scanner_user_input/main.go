package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // os.Stdin get the input from the user in command line
	fmt.Printf("%T\n", scanner)           // Scanner is a pointer to the bufio scanner

	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text:", text)
	fmt.Println("Input bytes:", bytes)

	// If we want to read the input from the file than we use "<" operator
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)
		if text == "exit" {
			fmt.Println("Exiting the Scanning")
			break // it will terminate the loop
		}
	}

	fmt.Println("Just a message after the for loop")
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
