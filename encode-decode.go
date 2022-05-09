package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

func ReverseSlice(s []string) []string {
	for i := len(s)/2 - 1; i >= 0; i-- {
		pos := len(s) - 1 - i
		s[i], s[pos] = s[pos], s[i]
	}
	return s
}

func sumSlice(s []int) int {
	sum := 0
	for _, item := range s {
		sum += item
	}
	return sum
}

func ConvertSliceToString(input []int) string {
	var output []string
	for _, item := range input {
		output = append(output, strconv.Itoa(item))
	}
	return strings.Join(output, ",")
}
