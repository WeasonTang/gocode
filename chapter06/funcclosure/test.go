package main

import (
	"strings"
	"fmt"
)

func makeSuffix(suffix string) func(string) string {
	return func (fileName string) string {
		if strings.HasSuffix(fileName,suffix) {
			return fileName
		} 
		return fileName + suffix
	}
}


func main() {
	//需求：编写一个函数，具体要求如下
	//1.编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg),并返回一个闭包
	//2.调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀，则返回 文件名.jpg。反之，则返回原文件名
	//3.要求使用闭包的方式完成
	//4.strings.HasSuffix, 返回一个bool值 判断是否有指定的后缀

	f := makeSuffix(".jpg")
	fmt.Println("dogName = ",f("dog.jpg"))
	fmt.Println("catName = ",f("cat"))
}