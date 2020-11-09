package main

import (
	"fmt"
	"time"
)

/*func TryReceive(c <-chan int) (data int,more,ok bool ){
	select {
	case data,more = <- c:
		return data,more,true
	case <- time.After(time.Duration(1)):
		return 0,true,false
	}
	//return data,more,true
}*/
/*func TryReceive(In <- chan int, OutA, OutB chan int)  {

	for data := range In {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}*/

func main()  {
	// channels creation
	chan1 := make(chan int)
	chan2 := make(chan int)

	// send value
	go func() {
		chan1 <- 2
		time.Sleep(8 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(2 * time.Second)
	}()

	for i := 0; i < 2; i++ { //condition is based on values
		select {
		case first := <-chan1: // receive value
			fmt.Printf("Value: %d\n", first)
		case second := <-chan2:
			fmt.Printf("Value: %d\n", second)
		}
	}
}
//func main(){
//	TryReceive()
//}