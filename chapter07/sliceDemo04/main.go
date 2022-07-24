package main

import (
	"fmt"
)

func main() {

	var slice []int = []int{11, 22, 33, 44, 55}
	var slice2 = make([]int, 10)
	copy(slice2, slice)
	slice[2] = 77
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)
}