package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx,cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:\t", ctx.Err())
	fmt.Println("Current number of goroutines", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * time.Duration(200))
				fmt.Println("Working", n)
			}
		}

	}()

	time.Sleep(time.Second*2)

	fmt.Println("error check 2:\t",ctx.Err())
	fmt.Println("current number of goroutines",runtime.NumGoroutine())
	fmt.Println("About to cancel the context")
	
	cancel()

	fmt.Println("error check 3:\t",ctx.Err())
	fmt.Println("current of number of goroutines",runtime.NumGoroutine())
}
