package main
import (
	"fmt"
)

func main() {

	// 控制台输入 姓名、地址、收入、是否通过
	var name string
	var age byte
	var salary float32
	var isPass bool

	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)

	// fmt.Println("请输入收入")
	// fmt.Scanln(&salary)

	// fmt.Println("请输入通过与否")
	// fmt.Scanln(&isPass)

	// fmt.Printf("%v 的年龄是%v 收入是%v 通过情况%v", 
	// 	name, age, salary, isPass)
 
	//方式2 fmt.Printf
	fmt.Println("亲输入姓名 年龄 收入 是否通过 并用空格隔开")
	fmt.Scanf("%s %d %f %t",&name, &age, &salary, &isPass)
	fmt.Printf("%v 的年龄是%v 收入是%v 通过情况%v",
		name, age, salary, isPass)
}