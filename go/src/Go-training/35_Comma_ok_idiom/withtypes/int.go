package withtypes

import "fmt"

// send values to the channels

func Sendint( e,o,q chan<- int) {

	for i := 0 ; i < 10 ; i++ {

		if i % 2 == 0 {
			e <- i
		}else {
			o <- i
		}
		
	}

	close(q)
	
}


// Recive values of channel


func Recieveint ( e,o,q <-chan int) {

	for {

		select {

		case v := <-e :
			fmt.Println("The even value is",v)
		
		case v := <-o :
			fmt.Println("The odd value is",v)
		case i, ok := <-q :

			if !ok {
				fmt.Println("The value of comma ok succed", i,ok)
				return
			}else {
				fmt.Println("The value of comma ok",ok)
			}
			
		}
	}
	
}
