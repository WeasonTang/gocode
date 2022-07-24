package main

import (
	"fmt"
)

func main() {
	heroes := make(map[string]string, 10)
	heroes["hero01"] = "莱拉"
	heroes["hero02"] = "史卡夫"
	heroes["hero03"] = "塔卡"
	fmt.Println(heroes)
	fmt.Printf("有%v个英雄\n", len(heroes))
	for k, v := range heroes {
		fmt.Printf("%v = %v\n", k, v)
	}

	//使用for-range 遍历value为map的map
	stuMap := make(map[string]map[string]string)
	stuMap["001"] = make(map[string]string,3)
	stuMap["001"]["name"] = "杰克"
	stuMap["001"]["sex"] = "男"
	stuMap["001"]["address"] = "美国"

	stuMap["002"] = make(map[string]string,3)
	stuMap["002"]["name"] = "罗丝"
	stuMap["002"]["sex"] = "女"
	stuMap["002"]["address"] = "英国"
	fmt.Println(stuMap)
	for k1, v1 := range stuMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1{
			fmt.Printf("\t %v=%v\n", k2, v2)
		}
	}
}