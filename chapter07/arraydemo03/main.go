package main

import (
	"fmt"
)

func main() {
	heroes := [...]string{"塔卡", "女枪", "莱拉"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i ,v)
	}

}