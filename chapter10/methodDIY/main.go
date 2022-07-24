package main

import(
	"fmt"
)

type Integer int

func (i Integer) println() {
	fmt.Println("i = ", i)
}

func (i *Integer) plusOne() {
	*i++
}

type Student struct {
	Name string
	Age int
}

//给Student实现String
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	// var i Integer = 7
	// i.println()
	// i.plusOne()
	// i.println()

	stu := Student{
		Name : "Harry",
		Age : 22,
	}
	fmt.Println(&stu)

}