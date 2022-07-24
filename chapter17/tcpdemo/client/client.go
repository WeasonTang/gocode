package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	//功能一：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		if strings.Trim(line, "\r\n") == "exit" {
			fmt.Printf("client退出")
			break
		}

		//再将line 发送给服务器
		_, errW := conn.Write([]byte(line))
		if errW != nil {
			fmt.Println("conn.Write err=", err)
		}
	}

}
