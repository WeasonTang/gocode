package main
import (
	"fmt"
)

func main() {

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你已经成年了，要为自己负责！")
	} else {
		fmt.Println("你还未成年，需要有监护人陪伴！")
	}

}