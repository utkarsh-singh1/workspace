package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t",ctx.Err())
	fmt.Printf("context type: \t%T\n",ctx)
	fmt.Println("------------------------")

	ctx, canceled := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t",ctx.Err())
	fmt.Printf("context type: \t%T\n",ctx)
	canceled()
	fmt.Println("------------------------")


	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t",ctx.Err())
	fmt.Printf("context type: \t%T\n",ctx)

	fmt.Println("------------------------")	
	
	/// Example for context


	ctx1, cancel := context.WithCancel(context.Background())

	fmt.Println("err1 here:",ctx1.Err())

	fmt.Println("Go routines:", runtime.NumGoroutine())

	fmt.Printf("context type:\t%T\n",ctx1)

	// launch goroutine

	go func() {


		n := 0

		for {

			select {
			case <-ctx1.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working ",n)
			}
		}
		
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("err2 check:",ctx1.Err())
	fmt.Println("go routines",runtime.NumGoroutine())


	fmt.Println("About to be cancel")
	cancel()
	fmt.Println("Concelled context")


	time.Sleep(time.Second * 2)
	fmt.Println("err3 check:",ctx1.Err())
	fmt.Println("go routines",runtime.NumGoroutine())
	
}
