package main

import (
	"fmt"
)

func main() {
	//map声明
	var a map[string]string
	a = make(map[string]string, 10)
	a["hero01"] = "莱拉"
	a["hero02"] = "塔卡"
	a["hero02"] = "隐狐"
	fmt.Println(a)
	fmt.Println("size=", len(a))
}