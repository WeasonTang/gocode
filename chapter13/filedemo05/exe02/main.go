package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	filePath := "d:/file/test.txt"
	//案例2：打开一个存在的文件，覆盖原来的内容
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666 )
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入的语句
	str := "hello, Append\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	
	//因为writer是带缓存的，需要调用Flush方法
	writer.Flush()
}