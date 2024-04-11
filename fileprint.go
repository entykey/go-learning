package main

import (
	"fmt"
	"io"
	"os"
	"os/user" // to taget "Users/user/Documents" directory on macOS
	"path/filepath"
)

func main() {
	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		return
	}

	// Construct the path to the Documents directory
	documentsDir := filepath.Join(currentUser.HomeDir, "Documents")
	filePath := filepath.Join(documentsDir, "example.txt") // Update with the file name

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buf := make([]byte, 1024)

	// Read and print the file content asynchronously
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		if n == 0 {
			break // Reached end of file
		}

		fmt.Print(string(buf[:n]))
	}

	fmt.Println() // Add a new line after printing the file content
}
