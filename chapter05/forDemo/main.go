package main
import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("i的值：%d\n", i)
	}

	j := 0
	for j <= 10 {
		fmt.Printf("j的值：%d\n", j)
		j++
	}

	k := 1
	for {
		if k < 10 {
			fmt.Printf("k的值 %d\n",k)
		} else {
			break
		}
		k++
	}
}