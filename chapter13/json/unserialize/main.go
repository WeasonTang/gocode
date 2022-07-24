package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Birthday string
	Sal float64
	Skill string
}

//将json反序列化成struct
func unserializeStruct() {
	str := "{\"name\":\"孙悟空\",\"age\":500,\"Birthday\":\"未知\",\"Sal\":2333,\"Skill\":\"养马\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", monster)
}

//将json反序列化成map
func unserializeMap() {
	str := "{\"age\":23,\"name\":\"索隆\",\"weapon\":\"三刀流\"}"
	var heroes map[string]interface{} 
	//反序列化 不用make, make的操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &heroes)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后 heroes=%v\n", heroes)
}

//将json反序列化成slice
func unserializeSlice() {
	str := "[{\"name\":\"良品铺子\",\"price\":99.9,\"size\":30}," + 
	"{\"name\":\"三只松鼠\",\"price\":69.9,\"size\":25}]"
	var slice []map[string]interface{} 
	//反序列化 不用make, make的操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	//struct反序列化
	unserializeStruct()
	//map反序列化
	unserializeMap()
	//slice反序列化
	unserializeSlice()
}