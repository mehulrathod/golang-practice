package main

import (
	"fmt"
	"strings"
)

func joinStr(element...string)string{
	return strings.Join(element, "-")
}

func main() {
	fmt.Println(joinStr())

	element := []string{"Mehul","Krishnakant","Rathod"}
	fmt.Println(joinStr(element...))

	fmt.Println(joinStr("Mehul","Rathod"))
	fmt.Println(joinStr("Mehul","Krishnakant","Rathod"))
}
