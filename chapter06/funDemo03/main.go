package main
import (
	"fmt"
)

//1. 在Go中，函数也是一种数据类型，可以赋值给一个变量，
//则该变量就是一个函数类型的变量，通过该变量可以对函数调用

func getSum(n1 int, n2 int) int{
	return n1 + n2
}

//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(diy func(int,int) int,num1 int, num2 int) int{

	return diy(num1, num2)
} 


func main() {

	a := getSum
	fmt.Printf("a的类型%T, getSum的类型是%T\n", a, getSum)
	res := a(10,60)
	fmt.Println("res=\n", res)

	//
	res2 := myFun(getSum, 77, 33)
	fmt.Println("res2=", res2)
}
