package main

import (
	"fmt"
	"reflect"
)

//反射演示
func reflect01(b interface{}) {

	//2.获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v type=%T\n", rValue, rValue)

	rValue.Elem().SetInt(20)
}

func main() {

	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
}
