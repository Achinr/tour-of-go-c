package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Printf("sum = %d \n", sum)
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("for 循环第 %d 次 \n ", i)
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
