package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func main() {
	//接口接受结构体时需为指针类型
	var i I
	i = &T{"Hello"}
	descibe(i)
	i.M()

	i = F(math.Pi)
	descibe(i)
	i.M()
}

func descibe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
