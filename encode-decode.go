package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeZone := "Asia/Shanghai"
	loc, _ := time.LoadLocation(timeZone)
	currentTime := time.Now().In(loc)
	fmt.Println(currentTime)

	fruits := []string{"Mango", "Grapes", "Kiwi", "Apple", "Grapes"}
	fmt.Println(RemoveDuplicates(fruits))
}

func RemoveDuplicates(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Shuffle(arrary []string) []string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arrary {
		j := random.Intn(i + 1)
		arrary[i], arrary[j] = arrary[j], arrary[i]
	}
	fmt.Println("Shuffle array:", arrary)
}
