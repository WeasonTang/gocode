package main

import (
	"fmt"
	"io"
	"net" //网络socket开发时，net包非常重要
)

func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close() //关闭连接

	for {
		//创建一个新切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		// fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("客户端%s 退出\n", conn.RemoteAddr().String())
			return
		}
		//3. 显式客户端发送的内容到服务端
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听了....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭

	//循环等待客户端连接
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}

	// fmt.Printf("listen suc=%v\n", listen)
}
