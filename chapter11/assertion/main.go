package main

import (
	"fmt"
)

//循环判断参数类型
func TypeJudge(items... interface{}) {
	for index, x := range items {
		switch x.(type) {
			case bool :
				fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
			case float32 :
				fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
			case float64 :
				fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
			case string :
				fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
			case Student :
				fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
			case *Student :
				fmt.Printf("第%v个参数是 Student指针 类型，值是%v\n", index, x)
			default :
				fmt.Printf("第%v个参数 类型不确定，值是%v\n", index, x)
		}

	}
}

type Student struct {
	Name string
}

func main() {
	var n1 int32 = 77
	var n2 float32 = 1.1
	var n3 float64 = 2.2
	var hero string = "索隆" 
	stu1 := Student{"黄桃"}
	stu2 := &Student{"罐头"}

	TypeJudge(n1, n2, n3, hero, stu1, stu2)
}