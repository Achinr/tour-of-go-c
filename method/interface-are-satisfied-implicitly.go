package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 类型通过实现一个接口的所有方法来实现该接口
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
