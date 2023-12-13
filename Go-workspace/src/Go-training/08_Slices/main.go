package main

import "fmt"

func main() {

	// Like array they store values but they are flexible in length

	// Slice literal

	xs := []int{1,2,3,4}

	a := []int{1,2,3,4}

	b := []int{0,0,0,0}

	fmt.Println(xs,a,b)

	fmt.Println(len(a),cap(a))


	for i:=0 ; i < len(xs) ; i++{
		b[i] = a[i] + xs[i]
	}

	fmt.Println(b)

	// Append to a slice

	xi := []int{32,34,56}

	fmt.Println(xi)

	xi = append(xi, 32,34,45,67,68)

	fmt.Println(xi)

	xi = append(xi, a...)

	fmt.Println(xi)

	// Slice a slice

	xi1 := []int{1,2,4,5,21,43,52,8,6,7,9}

	fmt.Println(xi1[0:3])
	fmt.Println(xi[3:9])
	fmt.Println(xi1[:10])
	fmt.Println(xi1[5:])
	fmt.Println(xi[:4])

	xi1 = append(xi1, 3,4,5)

	fmt.Println(xi1)


	// Delete a slice -> delete 43 from xi1

	// 43 index is 5

	xi1 = append(xi1[:5], xi1[6:]...)

	fmt.Println(xi1)


	// Use make to create a slice

	// syntax -> slice := make([]type, len, cap)

	si := make([]int,4,8)

	fmt.Println(si,"length is",len(si),"and capacity is",cap(si))

	si = append(si, xs...)

	fmt.Println(si,"length is",len(si),"and capacity is",cap(si))

	si = append(si, b...)

	fmt.Println(si,"length is",len(si),"and capacity is",cap(si))

	si = append(si[:4], )

	fmt.Println(si,"length is",len(si),"and capacity is",cap(si))


	// Multi-dimensional slice

	c := [][]int{a,b}

	fmt.Println(c)

	// Slice is build upon array, means it always point towards array whenever it needs values, so any slice made of that slice always gonna point towards same array

	arr := []int{1,2,3,4,5}

	brr := arr

	fmt.Println(arr,brr)

	
	arr[4] = 56

	fmt.Println(arr,brr)

	// Now what to do if any slice that gonna use slice arr value but not point towards that same undeline array, use copy()

	crr := make([]int, len(arr))

	copy(crr,arr)
	
	fmt.Println(arr,crr)

	
	arr[4] = 36

	fmt.Println(arr,crr)

	// Example

	xi2 :=  make([]int,0,10)

	xi2 = []int{2,4,6,5,23,44,34,56,77,91,8,10,18,19}

	x,y := findeven(xi2)

	fmt.Println(x,y,xi2)
	
}


// It will return number of even values in a slice/array 
func findeven(s []int) (int,[]int) {

	n := 0

	for k  := range s {
		if s[k] % 2 == 0 {
			n++
		}else{
			s[k] = s[k] + 1
		}
	}

	return n,s
}


// it will return number of odd values in a slice/array
// func findodd(s []int) int {

// }


