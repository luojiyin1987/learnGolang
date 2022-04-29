package main

import (
	"log"
	"os"
)

func main() {
	oldName := "text.txt"
	newName := "text2.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
