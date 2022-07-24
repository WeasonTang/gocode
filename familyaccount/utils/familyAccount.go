package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//声明一个变量，保存接收用户输入的选项
	key string
	//声明一个变量，控制是否退出for
	loop bool

	//定义账户的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支说明
	note string
	//定义个变，记录是否有收支的行为
	flag bool
	//收支详情 当有收支时，只需要对details 进行拼接处理即可
	details string
}

//方法0：工厂模式,返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount {
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说   明",
	}
}

//方法1：显式主菜单
func (this *FamilyAccount) MainMenu() {
	//显式主菜单
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("	   1 收支明细")
		fmt.Println("	   2 登记收入")
		fmt.Println("	   3 登记支出")
		fmt.Println("	   4 收退出软件")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()			
			default :
				fmt.Println("请输入正确的..")
		}
		if !this.loop {
			break
		}
	}
}
//方法2：显式收支明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录----------") 
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细， 请来一笔吧")
	}
} 

//方法3：收入
func (this *FamilyAccount) income() {
	fmt.Println("登记收入..")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money) //修改账户余额
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
} 

//方法4：支出
func (this *FamilyAccount) pay() {
	fmt.Println("登记支出..")
	fmt.Scanln(&this.money)
	//判断余额是否足够支出
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
} 

//方法5：退出
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出码? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}
