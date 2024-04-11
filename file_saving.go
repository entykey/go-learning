// STATUS: failed, downloaded file broken, won't be able to open.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url, filePath string) error {
	// Create the file to which the content will be written
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the file's metadata to determine its size
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a progress indicator
	fmt.Println("Downloading...")
	fileSize := resp.ContentLength
	progress := make(chan int64)
	go func() {
		var total int64
		for {
			n, ok := <-progress
			if !ok {
				break
			}
			total += n
			fmt.Printf("\rProgress: %.2f%%", float64(total)/float64(fileSize)*100)
		}
		fmt.Println("\nDownload complete!")
	}()

	// Copy the response body to the file while also updating the progress
	_, err = io.Copy(file, io.TeeReader(resp.Body, &progressReader{resp.Body, progress}))
	if err != nil {
		return err
	}

	return nil
}

type progressReader struct {
	reader   io.Reader
	progress chan<- int64
}

func (pr *progressReader) Read(p []byte) (n int, err error) {
	n, err = pr.reader.Read(p)
	pr.progress <- int64(n)
	return
}

// Implementing Write method for progressReader to satisfy io.Writer interface
func (pr *progressReader) Write(p []byte) (n int, err error) {
	n, err = pr.reader.Read(p)
	pr.progress <- int64(n)
	return
}

func main() {
	url := "https://images.dog.ceo/breeds/mountain-swiss/n02107574_1387.jpg" // Replace with the actual URL of the file
	fileName := "dog_image.jpg"

	// filePath := "dog_image.jpg"                                              // Replace with the name you want to save the file as
	// err := downloadFile(url, filePath)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("File downloaded successfully.")

	// Get the path to the Downloads directory
	downloadDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	downloadDir = filepath.Join(downloadDir, "Downloads")

	// Ensure the Downloads directory exists
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		err = os.Mkdir(downloadDir, 0755)
		if err != nil {
			fmt.Println("Error creating Downloads directory:", err)
			return
		}
	}

	// Construct the full file path
	filePath := filepath.Join(downloadDir, fileName)

	err = downloadFile(url, filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File downloaded successfully.")
}
