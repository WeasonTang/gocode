package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
	
)

//拷贝文件方法
func CopyFile(desFilePath string, srcFilePath string ) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("open srcFile err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcFile，获取 Reader
	reader := bufio.NewReader(srcFile)

	//打开desFilePath
	desFile, err := os.OpenFile(desFilePath, os.O_WRONLY | os.O_CREATE, 0666 )
	if err != nil {
		fmt.Printf("open desFile err=%v\n", err)
	}
	defer desFile.Close()
	//通过desFile获取writer
	writer := bufio.NewWriter(desFile)
	return io.Copy(writer, reader)

}


func main() {
	//调用CopyFile方法
	srcFilePath := "D:/图片/t01db0f61297ea50c01.jpg"
	desFilePath := "d:/file/pic.jpg"
	_, err := CopyFile(desFilePath, srcFilePath)
	if err == nil {
		fmt.Printf("拷贝成功\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}