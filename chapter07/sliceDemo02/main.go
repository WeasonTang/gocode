package main

import (
	"fmt"
)

func main() {
	//make创建切片
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[4] = 20
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	//第三种方式
	var sliceStr []string = []string{"水星","金星","地球","火星"}
	fmt.Println(sliceStr)
	fmt.Println("sliceStr的size=", len(sliceStr)) //4
	fmt.Println("sliceStr的cap=", cap(sliceStr)) //4

}