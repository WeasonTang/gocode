package main

import (
	"fmt"
)

func AddUpper() func (x int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(x+64)
		fmt.Println("str=", str)
		return n
	}
}

func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := AddUpper()
	fmt.Println(f2(1))
}
