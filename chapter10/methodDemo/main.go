package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//给Person类型绑定一份方法
func (person Person) test() {
	p.Name = "one"
	fmt.Println("我的名字是：", p.Name)
}


func main() {

	var p Person 
	p.Name = "zero"
	p.test() //调用方法
	fmt.Println("主函数name: ", p.Name)
}