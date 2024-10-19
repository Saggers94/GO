package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Pointer is a variable that stores a memory address of another variable
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	// Create function create a new file if it doesn't exist
	// If the file already exists it will be truncated
	// This function returns a file descriptor which is a pointer to os.File and an error value
	// newFile, err = os.Create("/etc/a.txt") // this will give an error because path is matching the linux path
	newFile, err = os.Create("a.txt")

	if err != nil {
		//Printing out the error if any
		// fmt.Println(err)
		// os.Exit(1)

		// Another way to print out the error
		// This is the idiomatic way to handle the error in GO
		// The preffered way to handle the error in GO
		log.Fatal(err) // log is thread safe and also adds a timing information automatically
	}

	// If we have 50 here, it means empty the file until 50 bytes and everything after
	// 50 bytes will be lost
	// err = os.Truncate("a.txt", 50)

	// 0 means Completely empty the file
	err = os.Truncate("a.txt", 0)
	if err != nil {
		log.Fatal(err)
	}

	// When we are done with the file we must close it
	newFile.Close()

	// Opening the file in read only mode
	// file, err := os.Open("a.txt")
	// we can also use the openFile function to open the file
	// Second argument is opening mode and the third argument is the permission
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644) // "|" is an OR operator
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	// Stat return the value of type fileInfo and an error value
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory:", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			// log.Fatal terminates the program so the next iteration of the code won't run
			// log.Fatal("File doesn't exist")
			// so changing it to the fmt.Println
			fmt.Println("File doesn't exist")
		}
		// log.Fatal(err)
	}

	// oldPath := "a.txt" // we have to be careful while defining the path because windows path and linux path are not compatible
	// newPath := "aaa.txt"
	// err = os.Rename(oldPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Removing the file
	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}
}
