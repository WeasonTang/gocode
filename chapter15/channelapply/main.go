package main

import (
	"fmt"
)
var running bool = true

//writeData
func writeData(intChan chan int, exitChan chan bool) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan<- i
		fmt.Printf("writeData 写入数据%v\n", i)
	}
	close(intChan)
	running = false;
}

//readData
func readData(intChan chan int, exitChan chan bool) {
	exitChan<- true
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// exitChan<- true
	running = false;
	close(exitChan)
}

func main() {

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan, exitChan)
	//go readData(intChan, exitChan)
	
	// _, ok := <-exitChan
	// fmt.Println("是否ok: ", ok)
	// _, ok2 := <-exitChan
	// fmt.Println("是否ok2: ", ok2)

	for {
		if !running {
			break
		}
	}
	
}