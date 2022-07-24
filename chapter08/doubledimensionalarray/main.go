package main

import (
	"fmt"
)

func main() {
	//使用方式一：先声明后赋值
	//声明二维数组  
	var arr[4][6]int
	//赋初值
	arr[1][2] = 1
	arr[2][3] = 7

	//使用方式二：直接初始化
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//遍历方式一
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	//遍历方式二
	for i, row := range arr2 {
		fmt.Printf("第%v行：", i)
		for _, column := range row {
			fmt.Print(column, " ")
		}
		fmt.Println()
	}
}