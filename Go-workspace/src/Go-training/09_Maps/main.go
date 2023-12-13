package main

import "fmt"

func main() {


	// syntax of a map -> map = map[sting]int{}

	// or

	// map := make(map[type]type1) 
	
	mp := map[int]bool{}

	mp[1] = true

	mp[2] = true

	mp[3] = false

	fmt.Println(mp)

	for _,value := range mp {
		fmt.Println(value)
	}

	mp1 := make(map[int](string))

	fmt.Printf("%T\n",mp1)


	// fmt.Print("Take a value from here ->")

	// fmt.Println(x,y)

	// x , _ := fmt.Scanln()
	
	if v,ok := mp[4] ; ok {
		fmt.Println(v,ok)
	}

	// Delete a key from map

	fmt.Println(mp[2])
	
	delete(mp, 2)

	fmt.Println(mp[2])

	
}
