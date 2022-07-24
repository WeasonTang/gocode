package main

import (
	"fmt"
)

func main() {

	//2.只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//num2 := <-chan2 //error
	close(chan2)

	//3.只读
	var chan3 <-chan int
	chan3 = make(chan int, 3)
	num3 := <-chan3
	fmt.Println("num3=", num3)
	// chan3 <- 30

}
