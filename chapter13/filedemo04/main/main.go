package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//案例1：创建一个新文件，写入5句从 "hello, Gardon"
	filePath := "d:/file/createTest.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666 )
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入的语句
	str := "hello, Gardon\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	
	//因为writer是带缓存的，需要调用Flush方法
	writer.Flush()
}