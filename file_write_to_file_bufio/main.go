package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// O_WRONLY : write only
	// O_CREATE : Create the file if it does not exist
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// defer will close the file after the complete execution of the main function
	defer file.Close()

	//Buffering before writing the bytes to the file
	// It is very useful if we want to do additonal operations before writing the bytes to the file
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99} // A:97 , B:98, C:99
	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)

	// Bytes available in buffer
	bytesAvailable := bufferedWriter.Available()
	// this will return 4093 because the size of the buffer is 4096 and 3 bytes are already written to the buffer
	log.Printf("Bytes available in buffer %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	// unflushed data that is waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	// this will return 24 because three bytes in the bytes slice and 21 bytes in the string
	// Byte has been Written to buffer but not to the file
	log.Printf("Bytes buffered %d\n", unflushedBufferSize)

	// Flush means flush the buffer to the file
	bufferedWriter.Flush() // write the buffer to the file

	// If we haven't written or flush the buffered data to the file that is waiting for the
	// flush and we want to reset the buffer than we can use the reset method
	bufferedWriter.Reset(bufferedWriter)
}
