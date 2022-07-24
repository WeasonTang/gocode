package main

import (
	"fmt"
)

//冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)
	var temp int
	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	} 
	fmt.Println("排序后arr=", *arr)
}

//二分查找
func BinaryFind(arr *[5]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//中间下标
	middle := (leftIndex + rightIndex)/2
	fmt.Println("middle = ", middle)
	if(*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else if(*arr)[middle] < findVal {
		BinaryFind(arr,middle + 1 , rightIndex, findVal)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}

func main() {
	//定义数组
	arr := [5]int{96, 24, 58, 77, 32}
	//调用排序函数
	BubbleSort(&arr)
	fmt.Println("main arr=", arr)
	BinaryFind(&arr, 0, len(arr) - 1, 32)
}