package main
import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("test hello,go " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}