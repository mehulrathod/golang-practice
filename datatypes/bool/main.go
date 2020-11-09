package main

import (
	"fmt"
	"unsafe"
)

const num = 10
/*Packages that import unsafe may be non-portable and are not protected by the Go 1 compatibility guidelines.*/
func main()  {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof(1 > num))
	fmt.Println(unsafe.Sizeof(!true))
	if num > 3 {
		fmt.Println(`"num" is greater than 3`)
	}
}


/*Default value of bool is false*/