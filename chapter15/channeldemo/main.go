package main

import (
	"fmt"
)

func main() {

	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.打印看看 intChan
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n",
		intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	<-intChan
	num := 211
	intChan <- num

	//4.看看管道长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n",
		len(intChan), cap(intChan))

	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n",
		len(intChan), cap(intChan))

}
