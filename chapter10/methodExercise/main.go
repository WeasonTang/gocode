package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (circle Circle) area() float64{
	circle.radius = 10
	return 3.14 * circle.radius * circle.radius
}

//为了提高效率，通常我们将方法和结构体的指针类型绑定
func (circle *Circle) area2() float64{
	var c3 Circle
	circle = &c3
	fmt.Printf("area2 实例的地址%p\n", circle)
	circle.radius = 10
	return 3.14 * (*circle).radius * (*circle).radius
}

func main() {

	// var c Circle
	// c.radius = 4.0

	// area := c.area()
	// fmt.Println("圆c的面积是：", area)
	// fmt.Println("圆c的半径是是：", c.radius)

	var c2 Circle
	c2.radius = 5.0
	fmt.Printf("main 实例的地址%p\n", &c2)

	area2 := c2.area2()
	fmt.Println("圆c2的面积是：", area2)
	fmt.Println("圆c2的半径是是：", c2.radius)
}