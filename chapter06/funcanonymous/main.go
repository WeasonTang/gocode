package main

import "fmt"

var (
	globalAnonymousFunc = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	//方式一：在定义匿名函数时就直接调用
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	//方式二：赋给变量
	a := func (n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10,40)
	fmt.Println("res2=", res2)

	res3 := globalAnonymousFunc(7,9)
	fmt.Println("res3=", res3)
}