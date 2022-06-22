package main

import (
	"fmt"
)

//更改切片的元素会修改其底层数组中对应的元素

func main() {
	names := [4]string{
		"john",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
