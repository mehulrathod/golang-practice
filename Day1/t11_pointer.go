package main

import "fmt"

/*
- If i have a variable as we've seen its very easy to output that but we can also get the memory address of the variable
by using ampersand and that gives us a pointer to variable
- If we have a function which incremented a variable and recalled but that's useless but if we pass pointer to the variable
as shown below then the function is going to be able to look at the value at that memory reference and modify the original
version
*/
func inc(x *int)  {
	*x++
}
/*
func inc(x int)  {
	x++
}
*/

func main()  {
	i := 7
	inc(&i)
	fmt.Println(i)
}

