package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileStat, err := os.Stat("emptyFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name", fileStat.Name())
	fmt.Println("Size", fileStat.Size())
	fmt.Println("Pesmissions", fileStat.Mode())
	fmt.Println("Last modified", fileStat.ModTime())
	fmt.Println("Is Directory", fileStat.IsDir())
}
