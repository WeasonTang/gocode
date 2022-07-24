package main
import (
	"fmt"
)

func main () {

	var key byte
	fmt.Println("请输入一个字符，a,b,c,d...")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一大概会很忙吧")
	case 'b':
		fmt.Println("星期二会显得我太心急不是吗")
	case 'c':
		fmt.Println("星期三有点模棱两可的感觉")
	case 'd':
		fmt.Println("星期四不知为何我就是不喜欢")
	case 'e':
		fmt.Println("喔~这周星期五")
	default:
		fmt.Println("mind control")
	}

}