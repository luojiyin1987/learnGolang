package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("end", s)
				ExitFunc()
			case syscall.SIGUSR1:
				fmt.Println("user1 ", s)
			case syscall.SIGUSR2:
				fmt.Println("user2", s)
			default:
				fmt.Println("default", s)
			}
		}
	}()

	fmt.Println("start")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func ExitFunc() {
	fmt.Println("start exit")
	fmt.Println("clear")
	fmt.Println("exit")
	os.Exit(0)
}
