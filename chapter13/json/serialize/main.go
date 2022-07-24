package main
import (
	"fmt"
	"encoding/json"
)

//结构体
type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Birthday string
	Sal float64
	Skill string
}
//将结构体monster序列化
func serializeStruct() {
	//演示
	monster := Monster{
		Name : "孙悟空",
		Age : 500,
		Birthday : "未知",
		Sal : 2333,
		Skill : "养马",	
	}
	//将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//将map序列化
func serializeMap() {
	//定义一个map
	var heroes map[string]interface{}
	heroes = make(map[string]interface{})
	heroes["name"] = "索隆"
	heroes["age"] = 23
	heroes["weapon"] = "三刀流"

	//序列化
	data, err := json.Marshal(heroes)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("heroes map序列化后=%v\n", string(data))
}

//将切片序列化
func serializeSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "良品铺子"
	m1["size"] = 30
	m1["price"] = 99.9
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "三只松鼠"
	m2["size"] = 25
	m2["price"] = 69.9
	slice = append(slice, m2)
	//序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice序列化后=%v\n", string(data))
}
func main() {
	//结构体序列化
	serializeStruct()
	//map序列化
	serializeMap()
	//切片序列化
	serializeSlice()
}

