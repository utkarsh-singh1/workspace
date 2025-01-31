package main

import (
	//	"context"
	"fmt"
	"sync"
	"time"
)

func main(){

	//c :=  make(chan int)

	var n int64 = 1

	var mu sync.Mutex

	var wg sync.WaitGroup

	wg.Add(100)
	
	go func() {
		for i := 0 ; i < 100 ; i++{
			mu.Lock()
			n = n + multiplyByInt(int64(i),n)
			fmt.Println("Value of n is",n)
			mu.Unlock()
			wg.Done()
		}
	}()

	wg.Wait()
}

func multiplyByInt(n,m int64) int64{
	time.Sleep(time.Duration(500)*time.Millisecond)
	return n*m
}
