package main

import "fmt"

func main() {

	// Method 1 describing array

	var arr [7]int

	arr[0] = 1

	fmt.Println(arr)

	val := 2
	
	for i:=1 ; i < len(arr) ; i++ {

		arr[i] = val

		val += 1
	}


	fmt.Println(arr)


	// Method 2 describing array

	// Array literal

	arr1 := [4]int{1,2,4,3}

	arr1[3] = 4


	// Metgod 3 describing array

	// Array literal -> Using ... in box to define a random size array , Complier determines the size of array

	arr2 := [...]int{1,2,3,4,5}

	fmt.Println(arr2)

	{
		// Create a different block and create a new scope

		arr3 := [...]string{"Hello","World","My New hope"}

		fmt.Println(arr3)
	}
	
	fmt.Println(arr2[3])

	fmt.Printf("Type of arr is %T , arr1 is %T, arr2 is %T \n", arr,arr1,arr2) // Type of an array is -> [len_arr]value_types

	fmt.Printf("Type of arr[2] is %T \n" , arr[2])
}
