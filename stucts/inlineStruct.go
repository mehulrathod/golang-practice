package main

import "fmt"

type blogPing struct {
	author string
	title string
	postId int
}

func main()  {
	b := struct {
		author string
		title string
		postId int
	}{
		author: "Mehul",
		title: "Let's Learn Golang",
		postId: 12345,
	}

	fmt.Println(b)
}
