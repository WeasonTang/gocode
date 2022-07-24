package main

import (
	"fmt"
	"time"
)

func main() {
	//go 中时间和日期相关函数和方法
	//1.获取当前时间 
	now := time.Now()
	fmt.Printf("now值：%v,类型：%T\n", now, now)
	//now值：2022-03-07 22:25:39.5714043 +0800 CST m=+0.003574401,类型：time.Time

	//2.通过now获取 年月日 时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	//3.格式化日期时间
	//方式一：使用Printf 或者 Sprintf
	dateStr := fmt.Sprintf("当前时间：%d-%d-%d %d:%d:%d\n",now.Year(), 
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v",dateStr)
	//方式二：Format根据layout指定的格式返回t代表的时间点的格式化文本表示
	//func (t Time) Format(layout string) string 各个数字固定
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("04")) //打印当前分钟
	fmt.Println()

	//4.时间常量和休眠
	//需求，每隔一秒钟打印一个数字，10退出
	for i := 0; i <= 10 ; i++ {
		fmt.Println("i = ", i)
		time.Sleep(1000 * time.Millisecond)
	}

	//5.获取当前unix时间戳 和 unixnano时间戳
	//func (t Time) Unix() int64 
	//func (t Time) UnixNano() int64 
	//作用：获取随机数
	fmt.Printf("unix时间戳=%v\n", now.Unix())
	fmt.Printf("unixnano时间戳=%v\n", now.UnixNano())

}