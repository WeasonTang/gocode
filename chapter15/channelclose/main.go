package main
import (
	"fmt"
)

func main() {

	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200
	//close 关闭管道
	close(intChan) 
	fmt.Println("channel intChan has closed")
// 	intChan<- 233 
	//channel intChan has closed panic: send on closed channel

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2<- i * 2 //放入100个数据到管道
	}

	//先关闭管道
	close(intChan2)

	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	// for len(intChan2) > 0 {
	// 	fmt.Println("v=", <-intChan2)
	// }
}