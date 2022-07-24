package cal
import (
	_ "fmt"
)

//被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(a int, b int) int{
	res := a - b
	return res
}