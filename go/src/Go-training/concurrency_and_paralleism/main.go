package main

import (
	"fmt"
	"runtime"
)

func main() {

	var os = runtime.GOOS
	var arch = runtime.GOARCH
	var cpus = runtime.NumCPU()
	var goroun = runtime.NumGoroutine()

	
	fmt.Println("OS\t\t",os)
	fmt.Println("Arch\t\t",arch)
	fmt.Println("CPUS\t\t",cpus)
	fmt.Println("Goroutine\t",goroun)



	go foo()
	bar()
	

	fmt.Println("Number of CPUS\t\t",cpus)
	fmt.Println("Number of Goroutines\t" ,goroun)

}

func foo() {

	fmt.Println(1)
}

func bar() {

	fmt.Println(2)
}

