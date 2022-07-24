package main
import (
	"fmt" 
	"unsafe"
)
func main() {
	var i int = 1
	fmt.Println("i=", i)

	var j int8 = -128
	fmt.Println("j= ",j)

	//测试一下 unit8的范围(0-255)
	var k uint8 = 255
	fmt.Println("k= ", k)

	//int, uint, rune, byte的使用
	var a int = -7777
	var b uint = 7777
	var c rune = -666
	var d byte = 255
	fmt.Println("a、b、c、d = ", a, b, c, d)

	//使用细节
	var n1 = 100 // ? n1 是什么类型 占用字节
	//1. fmt.Printf()  可以格式化输出
	//2. unsafe.Sizeof()  是unsafe包的一个函数，可以返回变量占用的字节数
	fmt.Printf("n1 的类型是：%T 占用的字节数是：%d \n", n1, unsafe.Sizeof(n1))
	//3.Go 程序中整型变量在使用时，遵守保小不保大的原则，
	//即：在保证程序正确运行下，尽量使用占用空间小的数据类型
	var age byte = 90
	fmt.Printf("age指针 %p",&age)
}