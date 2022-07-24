package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
	
)

//定义一个结构体，用于保存统计结果
type CharCount struct {
	EnCount int //记录英文个数
	NumCount int //记录数字的个数
	SpaceCount int //记录空格的个数
	OtherCount int //记录其它字符的个数
}

func main() {
	//思路：打开一个文件，创建一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	filePath := "d:/file/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer file.Close()
	//定义一个CharCount 实例
	var charCount CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环读取file内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历 str, 进行统计
		for _, v := range str {
			switch {
				case v >= 'a' && v <= 'z' :
						fallthrough //穿透
				case v >= 'A' && v <= 'Z' :
					charCount.EnCount++
				case v == ' ' || v== '\t' :
					charCount.SpaceCount++
				case v >= '0' && v <= '9' :
					charCount.NumCount++
				default :
					charCount.OtherCount++
			}
		}
	}
	//输出统计的结果
	fmt.Printf("英文字符的个数=%v\n 数字的个数=%v\n 空格的个数=%v\n 其它字符的个数=%v\n",
	charCount.EnCount, charCount.NumCount, charCount.SpaceCount, charCount.OtherCount)
}