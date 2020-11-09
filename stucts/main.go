package main

import "fmt"

type blogPost struct {
	author string
	title  string
	postId int
}

func main() {
/*	var b blogPost	// initialize the struct

	//fmt.Println(b)

	b = blogPost{
		author: "Mehul",
		title: "let understand go programming",
		postId: 1234,
	}
	b.author = "Akhil"

	fmt.Println(b.postId)
	fmt.Println(b.author)
	fmt.Println(b.title)*/
	b := blogPost{"Mehul", "Let's Learn about Golang", 123}
	fmt.Println(b)
}
