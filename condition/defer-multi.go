//推迟的函数调用会被压入一个栈中

package main

import "fmt"

func main() {
    fmt.Println("counting")
    
    for i := 10; i > 0; i-- {
        defer fmt.Println(i)
    }
    fmt.Println("done")
}
