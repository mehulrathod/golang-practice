package main

import "fmt"

func main()  {
//another way to loop
/*	i:=0
	for i<5 {
		fmt.Println(i)
		i++
	}
*/
	for i := 0; i<5; i++{
		fmt.Println(i)
	}
// another thing you can do with loop is iterate of each element in array or a slice by using range
	arr := []string{"a","b","c"}
	for index,value := range arr {
		fmt.Println("index:",index,"value:",value)
	}
	// same thing using map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key,value := range m {
		fmt.Println("key:",key,"value:",value)
	}

/*i := 1
fmt.Println(i)
i = i + 2
fmt.Println(i)
i = i + 2
fmt.Println(i)
i = i + 2
fmt.Println(i)
i = i + 2
fmt.Println(i)
i = i + 2
fmt.Println(i)*/

// 	for finding a odd number
/*	stoppingNumber := 30
	i := 1
	for i <= stoppingNumber {
		fmt.Println(i)
		i=i+2
	}
*/

// for printing a right half triangle
/*	m := 5
	for rows := 1; rows <= m; rows ++ {
		for col := 1; col <= rows; col++{
			fmt.Print("*")
		}
		fmt.Println()
	}
*/

// for printing a alphabet triangle
/*	for rows := 'A'; rows <= 'G'; rows ++ {
		for col := 'A'; col <= rows; col ++{
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
*/

// for creating triangle

}


//https://www.golangprograms.com/pyramid-star-series-and-patterns-programs-in-go-language.html