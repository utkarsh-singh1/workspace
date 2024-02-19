package main

import "fmt"

func main() {


	// Stack is self cleaning and efficient because whenever a function call is over stack cleans up the memory location that was occupied during memory allocation by values that are created during function calls

	x := 42

	//fmt.Println(x)

	// Now run the above code with "go run -gcflag -m main.go"

	// if Ouput of above command is -

	// x escapes to heap

	// x moves between main and fmt.Println


	// in case of a pointer return

	fmt.Println(&x)

	// command will print out moved to heap , means the value created during function call has escape the function and moved to a heap , a heap is a memory location where there is no automatic self-cleaning , garbage collector from go will clear out this memory stack in case it is not needed any more or else piled memory could shutoff system.

	
}
