package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a file
	f, err := os.Create("randomFile.txt")
	if err != nil {
		panic(err)
	}

	// Write to a file
	//size, err := file.Write([]byte("Hello World"))

	// Write string to a file instead of byte array
	size, err := f.WriteString("Hello Gopher!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created and wrote %d bytes in %s \n", size, f.Name())

	// Close a file
	f.Close()

	// Use os.Open followed by file.Read when you need more control over reading,
	// especially for larger files where you'd want to read in chunks or manage the file's lifecycle more closely.
	file, err := os.Open("randomFile.txt")
	if err != nil {
		panic(err)
	}

	// You can perform multiple Read operations on this file, potentially reading it in chunks using a buffer of a fixed size.
	reader := bufio.NewReader(file)
	buffer := make([]byte, 5)

	// Read the file in chunks of 5 bytes
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break // EOF
		}
		fmt.Printf("Read %d bytes from %s using bufio.NewReader method, content: %s\n", n, file.Name(), string(buffer[:n]))
	}
	// You have to remember to close the file after you're done with it, using the Close method. This can be a source of resource leaks if forgotten.
	file.Close()

	// Use os.ReadFile when you want to quickly read the entire contents of smaller files.
	// As it loads the entire file into memory, it might not be suitable for very large files.
	archive, err := os.ReadFile("randomFile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes from %s using ReadFile method\n", len(archive), "randomFile.txt")
	fmt.Println("File content: " + string(archive))

	// Delete a file
	err = os.Remove("randomFile.txt")
}
