package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("start")
	s := <-c
	fmt.Println("end", s)

}
