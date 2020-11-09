package main

import (
	"fmt"
	"time"
)

func form(s string) {
	for i:=0; i<3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	form("begin")

	go form("routine")

	go func (msg string) {
		fmt.Println(msg)
	}("routine between msg")

	time.Sleep(time.Second)
	fmt.Println("done")
}