package main

import "fmt"

type addPoint struct {
	author string
	title string
	postId int
}

func main()  {
	b := &addPoint{
		author: "Mehul",
		title: "Lets understand the blog",
		postId: 1245,
	}
	fmt.Println(*b)

	fmt.Println("Author's Name is :", b.author)
}