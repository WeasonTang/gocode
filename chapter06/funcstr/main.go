package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "hello嗨" 
	//1.统计字符串的长度，按字节 len(str) 
	fmt.Println("str len=",len(str)) //8 按字节返回

	//2.字符串遍历，同时处理有中文的问题 r := []rune(str)
	strArr := []rune(str)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("字符=%c\n",strArr[i])
	}

	//3.字符串转整数：func Atoi(s string) (i int, err error)
	n3, err := strconv.Atoi("233")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是：",n3)
	}

	//4.整数转字符串 func Itoa(i int) string
	str4 := strconv.Itoa(233)
	fmt.Printf("str4值：%v,类型：%T", str4, str4)

	//5.字符串转  []byte: var bytes = []byte("hello go") 
	var bytes = []byte("hello world")
	fmt.Printf("bytes=%v\n", bytes)

	//6.[]byte转 字符串：str = string([]byte{22,23,33})
	str6 := string([]byte{97, 98, 99})
	fmt.Printf("str6=%v\n", str6)

	//7. 10进制转 2,8,16 进制func FormatInt(i int64, base int) string
	str7 := strconv.FormatInt(77,2)
	fmt.Printf("77对应的二进制是%v\n",str7) 
	str7 = strconv.FormatInt(77,16)
	fmt.Printf("77对应十六进制是%v\n",str7) 

	//8.判断字符串s是否包含子串substr：func Contains(s, substr string) bool
	if strings.Contains("watermelo","water") {
		fmt.Println("watermelo has water")
	}

	//9.统计一个字符串有几个指定的子串：func Count(s, sep string) int
	num := strings.Count("banana","an")
	fmt.Printf("an个数%v\n", num)

	//10.不区分大小写的字符串比较(==区分字母大小写)：func EqualFold(s, t string) bool
	b1 := strings.EqualFold("abc","ABC")
	b2 := ("abc" == "ABC")
	fmt.Printf("b1 is %v, b2 is %v\n", b1, b2)

	//11.返回子串在字符串第一次出现的index,如果没有返回-1:func Index(s, sep string) int
	idx := strings.Index("ACC_qaq","qaq")
	fmt.Printf("idx=%v\n", idx)

	//12.返回子串在字符串最后一次出现的index,如果没有返回-1：func LastIndex(s, sep string) int
	lastIdx := strings.LastIndex("watemelo","me")
	fmt.Printf("lastIdx=%v\n",lastIdx)

	//13.将指定的子串替换成另外一个子串 n指替换个数,-1全替换：func Replace(s, old, new string, n int) string  
	str13 := "hello go go"
	str13New := strings.Replace(str13,"go","world",-1)
	fmt.Printf("新字符串：%v\n", str13New)

	//14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：func Split(s, sep string) []string
	str14Arr := strings.Split("banana","a")
	fmt.Printf("strArr=%v 类型： %T\n", str14Arr, str14Arr)

	//15.将字符串的字母进行大小写的转换：func ToLower(s string) string func | ToUpper(s string) string
	fmt.Println(strings.ToLower("GG")) 
	fmt.Println(strings.ToUpper("Gg"))

	//16.将字符串左右两边的空格去掉：func TrimSpace(s string) string
	fmt.Println(strings.TrimSpace(" wo yao xue go "))
	
	//17.返回将s前后端所有cutset包含的utf-8码值都去掉的字符串：func Trim(s string, cutset string) string
	fmt.Println(strings.Trim("!hello!","h!"))
	
	//18.返回将s前端所有cutset包含的utf-8码值都去掉的字符串：func TrimLeft(s string, cutset string) string
	fmt.Println(strings.TrimLeft("!hello!","!"))

	//19.返回将s后端所有cutset包含的utf-8码值都去掉的字符串：func TrimRight(s string, cutset string) string
	fmt.Println(strings.TrimRight("!hello!","!o"))

	//20.判断字符串是否以指定的字符串开头：func HasPrefix(s, prefix string) bool
	b20 := strings.HasPrefix("http://192.168.0.1","http")
	fmt.Printf("b20 is %v, type is %T\n",b20 ,b20)

	//21.判断字符串是否以指定字符串结束：func HasSuffix(s, suffix string) bool
	if strings.HasSuffix("main.go","go") {
		fmt.Println("main.go suffix is go")
	}
}