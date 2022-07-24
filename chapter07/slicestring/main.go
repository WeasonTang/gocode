package main 
import (
	"fmt"
)

func main() {

	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@qq.com"
	//使用切片获取到
	slice := str[6:]
	fmt.Println(slice)

//	str[0] = 'w'  报错(strings are immutable)
	slice2 := []byte(str)
	slice2[0] = 'w'
	str = string(slice2)
	fmt.Println("str=", str)

	slice3 := []rune(str)
	slice3[0] = '嗨'
	str = string(slice3)
	fmt.Println("str=", str)
}