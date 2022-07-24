package main
import (
	"fmt"
	"strconv"
)

func main() {

	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n",b ,b)

	var str2 string = "777777"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)

	var str3 string = "233.33"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//注意  如果没有转成功 用默认值
	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64) 
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)
}