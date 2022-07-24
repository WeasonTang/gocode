package main
import (
	"fmt"
)

func main() {
	var str = "hello world 哈哈"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

}