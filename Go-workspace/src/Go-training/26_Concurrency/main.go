package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t",runtime.GOARCH)
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("ROUTINES\t",runtime.NumGoroutine())

	wg.Add(1)
	// Adding go keyword creates another goroutine and send foo to run on that goroutine.(https://www.golang-book.com/books/intro/10#section1)
	go foo()
	bar()
	
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("ROUTINES\t",runtime.NumGoroutine())

	wg.Wait()
}


func foo() {

	for i :=0 ; i<10 ; i++  {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	for i := 0 ; i<8 ; i++ {
		fmt.Println(i)
	}
}
