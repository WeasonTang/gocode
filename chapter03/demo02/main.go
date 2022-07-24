package main
import "fmt"
func main() {
	//golang的变量使用方式1
	//第一种： 指定变量类型若不初始化会使用默认值
	var i int 
	fmt.Println("i=", i)

	//第二种： 类型推导
	var num = 7.7
	fmt.Println("num = ", num)

	//第三种 :=
	name := "weason"

}

