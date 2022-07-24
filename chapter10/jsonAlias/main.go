package main

import (
	"fmt"
	"encoding/json"
)

type Hero struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	hero := Hero{"阿卡丽", 22, "飞雷神"} 
	fmt.Println("hero", hero)

	jsonHero, err := json.Marshal(hero)
	if err != nil {
		fmt.Println("json处理错误：", err)
	}
	fmt.Println("jsonHero", string(jsonHero))
}