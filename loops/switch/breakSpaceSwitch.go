package main

import "fmt"

func main() {
Loop:
	for _, ch := range "z\na b\nc\nd f" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}
