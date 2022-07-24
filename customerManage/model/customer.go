package model

import (
	"fmt"
)

//声明一个customer结构体，表示一个客户信息

type Customer struct {
	id int
	name string
	gender string
	age int
	phone string
	email string
}

//编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer {
		id : id,
		name : name,
		gender : gender,
		age : age,
		phone : phone,
		email : email,
	}
}

//编写一个工厂模式，返回一个Customer的实例
func NewCustomerNoId(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer {
		name : name,
		gender : gender,
		age : age,
		phone : phone,
		email : email,
	}
}

//返回用户信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.id,
		this.name, this.gender, this.age, this.phone, this.email)
	return info		
}

func (this *Customer) SetId(id int) {
	this.id = id
}

func (this *Customer) GetId() int{
	return this.id
}

