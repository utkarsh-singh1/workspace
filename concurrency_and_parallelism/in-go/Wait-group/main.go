package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* As usual func main() is the entry point of all the functions and statement that is to be executed after running this file containing package main,

But before main() there is another function called as func init() that runs before main(), a single go file (even packages can have this) can have as many as it wants.


*/

var wg sync.WaitGroup


func main() {

	os := runtime.GOOS

	arch := runtime.GOARCH

	fmt.Println("OS and Architecure of this system is", os,"and" , arch)
	
	cpu := runtime.NumCPU()
	gorout := runtime.NumGoroutine()

	fmt.Println("Number of cpus and goroutines", cpu, gorout)

	foo()

	
	wg.Add(1)
	go bar()

	wg.Wait()
	fmt.Println("Number of Cpus and goroutines currently", cpu, gorout)

	

}

func foo() {

	for i := 0 ; i < 10 ; i++ {

		fmt.Println(i)
	}

	
	
}

func bar() {

	for j := 10 ; j < 20 ; j++ {

		fmt.Println(j)
	} 

	wg.Done()
}
