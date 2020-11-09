package main

import "fmt"


type Publisher interface {
	Publish()  error
}

type blogPost struct {
	author  string
	title   string
	postId  int
}


func (b blogPost) Publish() error {
	fmt.Printf("The title on %s has been published by %s\n" , b.title, b.author)
	return nil
}

// Receives any type that satisfies the Publisher interface
func PublishPost(publish Publisher) error {
	return publish.Publish()
}

func main() {

	var p Publisher

	fmt.Println(p)

	b := blogPost{"Alex","understand structs and interface types",12345}

	fmt.Println(b)

	PublishPost(b)

}