package main

import (
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("emptyFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File created successfully")
	emptyFile.Close()
}
