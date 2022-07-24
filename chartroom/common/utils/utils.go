package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

//读取数据
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8192)
	// fmt.Println("读取数据...")

	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		// fmt.Println("read pkg header error")
		return
	}
	//根据buf[:4] 转成一个 uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(transfer.Buf[0:4])

	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println(" conn.Read fail err=", err)
		return
	}
	//把buf反序列化成 message.Message
	json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("read pkg body error")
		return
	}
	// fmt.Println("读取数据成功 message=", mes)
	return
}

//发送数据
func (transfer *Transfer) WritePkg(data []byte) (err error) {

	//1 先发送data的长度
	//先获取到 data的长度->转成一个表示长度的byte切片
	var pkgLen uint32 = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkgLen)
	//发送长度
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// fmt.Printf("发送消息长度=%d\n 内容是%s\n", len(data), string(data))

	//2.发送消息本身
	n, err = transfer.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
	}
	// fmt.Println("发送消息成功 data=", string(data))
	return
}

func (transfer *Transfer) WriteMesPkg(mes *message.Message) (err error) {
	//序列化mes
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	transfer.WritePkg(data)
	return
}
