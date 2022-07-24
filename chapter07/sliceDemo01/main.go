package main 

import (
	"fmt"
)

func main() {
	var intArr = [...]int{7, 22, 11, 66, 88,99}
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice 的元素个数 =", len(slice))
	fmt.Println("slice 的容量 =", cap(slice))
	fmt.Println(intArr)
	fmt.Println(slice)
	
	fmt.Printf("intArr[1]地址%p\n", &intArr[1])
	fmt.Printf("slice[0]地址%p\n", &slice[0])

	slice[1] = 55
	fmt.Println(intArr)
	fmt.Println(slice)
}