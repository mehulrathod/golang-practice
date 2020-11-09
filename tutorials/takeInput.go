package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	myname string
	myStoreName string
)

func main()  {
/*	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Full Name : ")
	myname, _ := reader.ReadString('\n')
	fmt.Println(myname)*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Name: ")
	myname, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Rate out dosa center between 1 to 5 : ")
	myStoreName, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myStoreName), 64)

	fmt.Printf("Hello %v, Thanks for rating us with %v and Your rating is recorded at %v.\n", myname,myNumRating,time.Now().Format(time.Stamp))

	if myNumRating == 5 {
		fmt.Println("Bonus for team")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("We are improving our services")
	} else if myNumRating < 3{
		fmt.Println("We serious work on our side")
	} else {
		fmt.Println("Please rate our services")
	}
}
