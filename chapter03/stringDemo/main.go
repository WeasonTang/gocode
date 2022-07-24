package main
import (
	"fmt"
	_ "unsafe"
)

func main() {
	//string的基本使用
	var address string = "加油 会好起来的！"
	fmt.Println(address)

	//go 中字符串 注意事项和细节
	//2.字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的
	var str = "hello"
	//str[0] = 'a' 
	fmt.Println(str)

	//3.字符串的两种表示形式
	// 1).双引号，会识别转义字符
	str2 := "abc\nabc"
	fmt.Println(str2)
	
	// 2).使用反引号
	str3 := `
	package main
	import "fmt"
	func main() {
		var i = 1
		var j = 2
		var r = i + j
		fmt.Println("r ", r)

		var firstName = "Harry"
		var lastName = "Potter"
		var fullName = firstName + " " + lastName
		fmt.Println("fullName:",fullName)
	}
	`
	fmt.Println(str3)

	//4.字符串拼接方式
	var str4 = "hello" + " world"
	str4 += " haha!"
	fmt.Println(str4)

	//5.字符串拼接换行 + 号留上面
	 str5 := "hello " + "hello " + "hello " + "hello " + "hello " + "hello " + "hello " + 
	"hello " + "hello " + "hello " + "hello " + "hello " + "hello " + 
	"hello " + "hello " + "hello " + "hello " + "hello " 
	fmt.Println(str5)
}