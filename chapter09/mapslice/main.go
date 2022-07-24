package main

import (
	"fmt"
)

func main() {
	//长度可以动态变化的map切片

	//1.声明一个map切片
	monsters := make([]map[string]string, 2)
	//2. 增加第一个monster信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "金角"
		monsters[0]["age"] = "567"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "银角"
		monsters[1]["age"] = "555"
	}
	//通过内置函数append 动态增加切片
	newMonster := map[string]string{
		"name" : "灰犀牛",
		"age" : "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}