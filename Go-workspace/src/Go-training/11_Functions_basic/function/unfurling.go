package function

func Unfurling_sum(ui ...int) int {

	sum := 0

	for _,j := range ui {
		sum += j
	}

	return sum
}

func Unfurling_max(ui []int) int {

	max := 0

	for _,j := range ui {

		if j > max {
			max = j
		}
	}

	return max
}
