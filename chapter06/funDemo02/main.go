package main
import (
	"fmt"
)

func test01(n1 int) {
	n1 = n1 + 10
	fmt.Println("test01 n1 = ", n1)
}

func test02(n2 *int) {
	*n2 = *n2 + 20
	fmt.Printf("test02 *n2=%d n2=%v,address=%v\n ", *n2, n2, &n2)
}

func main() {

	//test 基本数据类型和数组默认都是值传递的，即进行值拷贝。
	//在函数内修改，不会影响到原来值
	n1 := 20
	test01(n1)
	fmt.Println("main n1 = ", n1)

	n2 := 7
	test02(&n2)
	fmt.Printf("main n2=%d,address=%v", n2, &n2)

}