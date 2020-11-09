package main

import "fmt"

func main(){
	sum := 0 //1,3,6,10
	for i := 1; i < 5; i++ {
		fmt.Println("=-=-=-=before-=-=-=-=-",i)
		fmt.Println("=-=-=-=before-=-=-=-=-",sum)
		sum += i //1,3,6,10
		fmt.Println("=-=-=-=inside-=-=-=-=-",sum)
	}
	fmt.Println("=-=-=-=outside-=-=-=-=-",sum) //10
}
/*
step-1:
	i = 1 , sum = 0
	o/p > 1
step-2:
	i = 2 , sum = 1
	o/p > 3
step-3:
	i = 3 , sum = 3
	o/p > 6
step-4:
	i = 4 , sum = 6
	o/p > 10
*/