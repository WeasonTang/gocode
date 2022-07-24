package main

import (
	"fmt"
	"errors"
)

func test() int{
	//使用defer + recover 来捕获并处理异常
	defer func() {
		//recover()内置函数，可以捕获异常
		err := recover() 
		if err != nil {
			fmt.Println("err=", err)
			//发送邮件给管理员
			fmt.Println("发送邮件给admin!qq.com")
		}
	}()
	n1 := 10
	n2 := 0
	res := n1/n2
	return res
}

//函数：取读取配置文件 init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func testRead() {
	err := readConf("config.int")
	if err != nil {
		panic(err)
	}
	fmt.Println("testRead后面的代码")
}


func main() {

	// res := test()
	// fmt.Printf("res=%v\n", res)
	testRead();
	fmt.Println("main()下面的代码..")
}