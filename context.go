package main

import (
	"context"
	"fmt"
	"time"
)

func timeoutHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doSth(ctx)

	time.Sleep(4 * time.Second)
	cancel()
}

func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds:\n", i)
		}
		i++
	}
}

func main() {
	fmt.Println("start")
	timeoutHandler()
	fmt.Println("end")
}
