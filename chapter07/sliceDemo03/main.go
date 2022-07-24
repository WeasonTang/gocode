package main
import (
	"fmt"
)
func main() {
	//方式一：使用常规的for循环遍历切片
	var arr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := arr[1:4]
	slice2 := slice[1:4]
	arr[2] = 99
	arr[4] = 77
	slice = append(slice, slice...)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%v\n", i, slice[i])
	}
	//方式二：for-range
	for i, v := range slice2 {
		fmt.Printf("slice2[%d]=%v\n", i, v)
	}
}