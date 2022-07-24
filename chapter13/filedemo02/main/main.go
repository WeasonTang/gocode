package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func main() {

	//打开文件
	//file叫法： 1.file对象 2.file指针 3.file文件句柄
	file, err := os.Open("d:/file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏
	
	//创建一个 *Reader,带缓冲
	/*
	const (
		defaultBufSize = 4096 //默认的缓冲区为4096
	)
	*/
	reader := bufio.NewReader(file)
	for {
		//读到一个换行就结束一次
		str, err := reader.ReadString('\n')
		//输出内容
		fmt.Print(str)
		// io.EOF 表示文件末尾
		if err == io.EOF { 
			fmt.Println()
			break
		}
	}
	fmt.Println("文件读取完毕")
}