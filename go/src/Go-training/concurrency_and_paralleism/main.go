package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var os = runtime.GOOS
	var arch = runtime.GOARCH
	var cpus = runtime.NumCPU()
	var goroun = runtime.NumGoroutine()

	fmt.Println("OS\t\t", os)
	fmt.Println("Arch\t\t", arch)
	fmt.Println("Intial CPUS\t\t", cpus)
	fmt.Println("Intial Goroutine\t", goroun)

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Number of CPUS\t\t", cpus)
	fmt.Println("Number of Goroutines\t", goroun)

	wg.Wait()
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
