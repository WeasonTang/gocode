package main
import (
	"fmt"
)
func main() {
	num1 := 100
	fmt.Printf("num1的类型%T, num1的值=%v, num1的地址%v\n",
		num1, num1, num1)
	//2.内建函数new分配内存。其第一个实参为类型，而非值。
	//其返回值为指向该类型的新分配的零值的指针。
	//func new(Type) *Type
	num2 := new(int)
	*num2 = 77
	//num2的类型*int, num2的值=0xc0000aa058(系统分配),
	// num2的地址0xc0000ce020(系统分配), num2指向的值=77
	fmt.Printf("num2的类型%T, num2的值=%v, num2的地址%v, num2指向的值=%v\n",
		num2, num2, &num2, *num2)
}