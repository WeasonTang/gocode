package main
import (
	"fmt"
)

func main() {
	var intArr [3]int
	//当我们定义完数组后，数组各元素有默认值0
	fmt.Println(intArr)

	fmt.Printf("intArr的地址=%p\n", &intArr)
	fmt.Printf("intArr[0]的地址=%p\n", &intArr[0])
	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])

	//从终端循环输入5个成绩，保存到float64数组,并输出
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素的值\n", i + 1)
	// 	fmt.Scanln(&score[i])
	// }

	// //遍历数组打印
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d]=%v\n", i, score[i])
	// }

	//四种初始化数组方式
	var arr1 [3]int = [3]int{1, 2, 4}
	fmt.Println("arr1=", arr1)

	var arr2 = [3]int{3, 6, 9}
	fmt.Println("arr2=", arr2)

	var arr3 = [...]int{2,3,3}
	fmt.Println("arr3=", arr3)

	var arr4 = [...]int{1: 233, 0: 777, 4: 666}
	fmt.Println("arr3=", arr4)

}