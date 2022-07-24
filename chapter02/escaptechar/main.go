package main

import "fmt" //fmt包中提供格式化 输出，输入的函数
func main() {
	//演示转义字符的使用  \t
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("c:\\user\\Adminstrator")
	fmt.Println("老板说\"明天放假！\"")
	// \r 回车，从当前行的最前面考试输出，覆盖掉以前的内容
	fmt.Println("人生若只如初见\r纳兰")
}
