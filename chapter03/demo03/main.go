package main
import "fmt"
//全局变量
var (
	time = 2021126
	address = "武汉"
	person = "卡夫卡"
)


func main() {

	//一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//方式二
	var name, age, sex = "指针", 23, "未知"
	fmt.Println("name=", name, "age=", age, "sex=", sex )

	//凡是三
	fruit,animal,count := "lemon","giraffe",77
	fmt.Println(fruit,animal,count)

	fmt.Println(time, address, person)
}