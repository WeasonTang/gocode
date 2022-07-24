package main

import (
	"fmt"
	"reflect"
)

//反射演示
func reflectTest01(b interface{}) {

	//通过反射获取传入变量的 type, kind, 值
	//1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v type=%T\n", rType, rType)

	//2.获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v type=%T\n", rValue, rValue)

	//3.将 rValue 转成 interface{}
	iv := rValue.Interface()
	fmt.Printf("iv=%v type=%T\n", iv, iv)
}

func reflectTest02(b interface{}) {
	//通过反射获取传入变量的 type, kind, 值
	//1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v type=%T\n", rType, rType)

	//2.获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v type=%T\n", rValue, rValue)

	//3.将 rValue 转成 interface{}
	iv := rValue.Interface()
	fmt.Printf("iv=%v type=%T\n", iv, iv)
	//断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("name=%v\n", stu.Name)
	}

	//4.获取变量对应的Kind
	// rKind := rType.Kind()
	rKind := rValue.Kind()
	fmt.Printf("rKind=%v, type=%T\n", rKind, rKind)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//int类型
	num := 100
	reflectTest01(num)

	stu := Student{
		Name: "奇犽",
		Age:  20,
	}

	reflectTest02(stu)
}
