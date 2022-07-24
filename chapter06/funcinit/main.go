package main
import (
	"fmt"
	//引入utils包
	"gocode/chapter06/funcinit/utils"
)

var age = getAge()

//验证全局变量是先于init初始化
func getAge() int {
	fmt.Println("获取age")
	return 28;
}


//init函数 通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main....age =", age)
	fmt.Println("utils_age = ", utils.Age, "name =", utils.Name)
}