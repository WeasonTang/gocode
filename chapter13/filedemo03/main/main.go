package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//没有显式的open文件，也不需要显式的close文件
	//因为，文件的open和close被封装到ReadFile函数内部
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/file/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err=%v", err)
	}
	//把读取到的内容显式在终端
	fmt.Printf("%v", string(content))

}