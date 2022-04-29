package main

import (
	"io"
	"log"
	"os"
)

func main() {
	sourceFile, err := os.Open("emptyFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create("emptyFile2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %d bytes.", bytesCopied)
}
