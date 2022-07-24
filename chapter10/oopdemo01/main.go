package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func(stu *Student) say() string {
	infostr := fmt.Sprintf("学生的信息 name=[%v], gender=[%v] age=[%v] id=[%v] score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
		return infostr
}	

type Box struct {
	length float64
	width float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.length * box.width * box.height
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元\n", visitor.Name, visitor.Age)		
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 不收费\n", visitor.Name, visitor.Age)		
	}
}


func main() {

	//测试
	var stu = Student{
		name : "L",
		gender : "male",
		age : 16,
		score : 99999.99,
	}
	stuInfo := stu.say()
	fmt.Println(stuInfo)

	box := Box{
		length : 2.5,
		width : 2.0,
		height : 1,
	}
	fmt.Println(box.getVolumn())

	var v Visitor

	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "exit" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}