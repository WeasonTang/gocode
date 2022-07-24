package main

import (
	"fmt"
	"gocode/chapter10/factory/model"
)

func main() {

	stu := model.NewStudent("kkk", 77.7)
	fmt.Printf("%p\n", stu)
	fmt.Printf("%v\n", &stu)
	fmt.Printf("Name[%v] Score[%v]\n", stu.Name, stu.GetScore())
}
