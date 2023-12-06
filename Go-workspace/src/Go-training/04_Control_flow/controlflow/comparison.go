package controlflow

import "fmt"

func GreaterThen(x,y int) string {
	if x > y {
		s := fmt.Sprintf("%v is greater than %v",x,y)
		return s
	}else if x == y {
		return "Every one is Equal my broda"
	}

	return "Unfortunately it is not in my domain to do it any lesser work, You can try my another counterpart Lessthan()"
}

func LessThen(x,y int) string {
	if x < y {
		s := fmt.Sprintf("%v is less than %v",x,y)
		return s
	}else if x == y {
		return "Every one is Equal my broda"
	}

	return "Unfortunately it is not in my domain to do it any greater work, You can try my another counterpart GreaterThan()"
}
