package function

func Variadic_sum ( ii ...int) int {

	sum := 0

	for _,j := range ii {
		sum += j
	}

	return sum
}


func Variadic_max (ii ...int) int {

	max := 0

	for _,j := range ii {
		if j > max {
			max = j
		}
	}

	return max
}

func Variadic_min (ii ...int) int {

	min := ii[0]

	for _,j := range ii[1:] {
		if min > j {
			min = j
		}
	}

	return min
}


