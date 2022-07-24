package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {

	for i := 2; i < 80000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
			}
		}

		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum 协程退出...")
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)

	//表示退出
	exitChan := make(chan bool, 8)

	start := time.Now().UnixNano()

	//开启一个协程，向 intChan放入 1-8000
	go putNum(intChan)
	//开启多个协程，从intChan取出数据，并判断是否为素数
	//是，则放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//primeChan 读取判断
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		end := time.Now().UnixNano()
		fmt.Printf("耗时：%d\n", end-start)
		close(primeChan)
	}()

	total := 0
	for v := range primeChan {
		total += v
	}
	fmt.Printf("素数和=%d\n", total)
	fmt.Println("主线程执行完毕")

}
