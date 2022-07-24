package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func (phone Phone) Start() {
	fmt.Println("手机开始工作...")
} 

func (camera Camera) Start() {
	fmt.Println("相机开始工作...")
} 

func (phone Phone) Stop() {
	fmt.Println("手机停止工作...")
} 

func (camera Camera) Stop() {
	fmt.Println("相机停止工作...")
} 

type Computer struct {}

func (computer Computer) working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	phone := Phone {
		Name : "一加",
	}
	camera := Camera {
		Name : "相机",
	}
	computer := Computer{}
	//1. 多态参数
	computer.working(phone)
	computer.working(camera)

	//2.多态数组
	var usbArr[3]Usb
	usbArr[0] = Phone{"苹果"}
	usbArr[1] = phone
	usbArr[2] = camera

	fmt.Println(usbArr)
}