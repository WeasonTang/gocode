package main

import (
	"fmt"
)

func main() {

	heroes := make(map[string]string, 10)
	heroes["hero01"] = "莱拉"
	heroes["hero02"] = "史卡夫"
	fmt.Println(heroes)

	//增加和修改
	heroes["hero03"] = "塔卡"	
	fmt.Println(heroes)
	heroes["hero03"] = "隐狐"	
	fmt.Println(heroes)

	//查找
	hero, exist := heroes["hero01"]
	if exist {
		fmt.Printf("hero %v is exist\n", hero)
	} else {
		fmt.Println("hero is not exist")
	}

	//删除
	delete(heroes, "hero02")
	fmt.Println(heroes)
	//如果想一次性删除所有entry,
	//1.遍历逐个删除 
	//2.直接make一个新空间
	heroes = make(map[string]string)
	fmt.Println(heroes)
}