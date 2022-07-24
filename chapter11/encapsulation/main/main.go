package main

import (
	"fmt"
	"gocode/chapter11/encapsulation/model"
)

func main() {
	p := model.NewPerson("jack")
	p.SetAge(23)
	p.SetSal(2333)
	fmt.Println(p)
	fmt.Println(p.Name, "age = ", p.GetAge(), "sal = ", p.GetSal())
}