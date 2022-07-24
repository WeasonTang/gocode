package main
import (
	"fmt"
	"gocode/chapter06/utils"
)

func main() {
	fmt.Println("hello world")
	res := utils.Cal(1.2, 2.3, '*')
	fmt.Println("res= ", res)
}