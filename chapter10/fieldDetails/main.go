package main

import (
	"fmt"
)

//如果结构体的字段类型是：指针，slice,和map的零值都是nil,即还没有分配空间
//如果需要使用这样的字段，需要先make,才能使用
type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int  //指针
	slice []int //切片
	map1 map[string]string //map
}


func main() {

	//定义一个结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.map1 == nil {
		fmt.Println("ok")
	}

	//使用slice 切片直接赋值会报错，需先make
	// p1.slice[0] = 100  
	p1.slice = make([]int, 5)
	p1.slice[0] = 100 
	fmt.Println(p1)

	//使用map map直接赋值会报错，需先make
	// p1.map1["key1"] = "hehe" 
	p1.map1 = make(map[string]string)
	p1.map["key1"] = "haha"
	fmt.Println(p1)

}