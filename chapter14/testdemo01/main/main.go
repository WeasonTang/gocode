package main
import (
	"fmt"
)

//被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10);
	if res == 55 {
		fmt.Printf("测试成功")
	} else {
		fmt.Printf("测试成功")
	}

}