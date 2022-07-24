package main

import (
	"fmt"
)

func main() {
	//方式一: 先声明再make
	var heros map[string]string
	heros = make(map[string]string, 10)
	heros["hero01"] = "莱拉"
	heros["hero02"] = "塔卡"
	heros["hero03"] = "格温"
	fmt.Println(heros)

	//方式二：声明时make
	cities := make(map[string]string, 10)
	cities["c01"] = "荆州"
	cities["c02"] = "武汉"
	cities["c03"] = "上海"
	fmt.Println(cities)

	//方式三：
	languages := map[string]string{
		"lan01" : "English",
		"lan02" : "中文",
	}
	languages["lan03"] = "Go"
	fmt.Println(languages)

	//案例演示
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
}