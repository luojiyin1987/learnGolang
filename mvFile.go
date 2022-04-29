package main

import (
	"log"
	"os"
)

func main() {
	olLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.text"

	err := os.Rename(olLocation, newLocation)

	if err != nil {
		log.Fatal(err)
	}
}
