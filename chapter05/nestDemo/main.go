package main
import (
	"fmt"
)

func main() {

	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)

	if second <= 8 {
		//进入决赛
		var gender string
		fmt.Println("请输入性别：男 或 女")
		fmt.Scanln(&gender)
		if gender == "女" {
			fmt.Println("恭喜你，进入女子组！")
		} else if gender == "男" {
			fmt.Println("恭喜你，进入男子组！")
		} else {
			fmt.Println("请输入正确的性别！")

		}
	} else {
		fmt.Println("很遗憾，你没有进入决赛！")
	}

}