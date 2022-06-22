package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	descibe(i)
	i.M()
}

func descibe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
