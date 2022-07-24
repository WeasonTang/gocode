package main
import (
	"fmt"
	"os"
)

func main() {

	//打开文件
	//file叫法： 1.file对象 2.file指针 3.file文件句柄
	file, err := os.Open("d:/file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出文件
	fmt.Printf("file=%v", file)
	fmt.Println()

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}