package main
import (
	"fmt"
)

func main() {

	var i = 10
	fmt.Println("i的地址 = ",&i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n",*ptr)

	var num int
	fmt.Printf("num的地址值%p\n", &num)
	fmt.Printf("num修改前的值%d\n",num) 
	
	var ptrn *int = &num
	*ptrn = 777
	fmt.Printf("num修改后的值%d\n",num) 

}