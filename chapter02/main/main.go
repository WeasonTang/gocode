package main

import "fmt"

// type student struct {
// 	name string
// 	age  int
// }

type name string

func (n name) len() int {

	return len(n)

}

func main() {

	// var x int = 10
	// fmt.Println(x)

	// p := &x
	// fmt.Printf("%v\n", p)
	// fmt.Printf("%x\n", p)
	// fmt.Printf("%p\n", p)
	// fmt.Println(p)

	// *p = 20
	// fmt.Println(*p)

	// var stu = student{
	// 	name: "tom",
	// 	age:  23,
	// }
	// fmt.Println(stu)
	// var stup *student = &stu
	// fmt.Printf("stu地址%p\n", &stu)
	// fmt.Printf("stu地址%p\n", stup)
	// fmt.Printf("stu地址%x\n", stup)
	// fmt.Printf("stu地址%v\n", stup)
	var myname name = "taozs嗨" //其实就是字符串类型

	l := []byte(myname) //字符串转字节数组

	fmt.Println(len(l)) //字节长度

	fmt.Println(myname.len()) //调用对象的方法

}
