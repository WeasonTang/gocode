package main

import (
	"fmt"
	"gocode/customerManage/service"
	"gocode/customerManage/model"
)

type customerView struct {
	//定义必要的字段
	key string //接收用户输入
	loop bool //表示是否循环的显式主菜单
	customerService *service.CustomerService
}

//显式客户列表
func (this *customerView) list() {
	//首先，获取当前所有的客户信息(切片中)
	customers := this.customerService.List()
	//显式客户信息
	fmt.Println("----------客户信息管理软件----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i< len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n----------客户列表打印完成----------\n\n")
}

//新增客户
func (this *customerView) add() {
	fmt.Println("----------添加客户----------")	
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	//构建没有id的customer实例
	customer := model.NewCustomerNoId(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("----------添加完成----------")	
	} else {
		fmt.Println("----------添加失败----------")	
	}
}

//删除客户
func (this *customerView) delete() {
	fmt.Println("----------删除客户----------")	
	fmt.Println("请选择待删除的客户编号(-1退出)")	
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("----------删除完成----------")	
		} else {
			fmt.Println("----------删除吃白用户编号不存在---")	
		}
	}
}

//退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否删除(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("输入有误，请输入Y/N")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}


//显式主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n----------客户信息管理软件----------")
		fmt.Println("	   1 添加客户")
		fmt.Println("	   2 修改客户")
		fmt.Println("	   3 删除客户")
		fmt.Println("	   4 客户列表")
		fmt.Println("	   5 退出软件")
		fmt.Print("请选择(1-5): ")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1" :
				this.add()
			case "2" :
				fmt.Println("修改客户")
			case "3" :
				this.delete()
			case "4" :
				this.list()
			case "5" :
				this.exit()
			default :
				fmt.Println("输入有误，请重新输入。。。")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理系统")
}

func main() {
	//在主函数中创建一个customerView实例
	customerView := customerView {
		key : "",
		loop : true,
		customerService : service.NewCustomerService(),
	}
	//显式主菜单
	customerView.mainMenu()
}
