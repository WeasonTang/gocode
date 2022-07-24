package main

import (
	"fmt"
)

//定义cat结构体，将cat字段放入
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}


func main() {
	//创建Cat实例
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃<。)#)))≦"
	fmt.Println(cat1)
	fmt.Println("Name=",cat1.Name)
	fmt.Println("Age=",cat1.Age)
	fmt.Println("Color=",cat1.Color)
	fmt.Println("Hobby=",cat1.Hobby)
}
