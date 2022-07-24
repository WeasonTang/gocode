package main
import (
	"fmt"
)

func main() {

	var price float32 = 89.23
	fmt.Println("price= ",price)

	//浮点数 = 符号位 + 指数位 + 尾数位
	//尾数部分可能丢失，造成精度损失
	var num1 float32 = -233.0000709
	var num2 float64 = -233.0000709
	fmt.Println("num1=", num1, "num2=", num2)

	//go 浮点型默认为float64 
	var num3 = 1.2
	fmt.Printf("num3的数据类型是 %T \n", num3)


}