package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("conn err")
		return
	}

	data := []byte("hello")

	//客户端发送数据
	conn.Write(data)

	//接受服务端的返回
	buf := make([]byte, 1024)

	cnt, err := conn.Read(buf)
	fmt.Printf("服务端数据长度:%d, 数据:%s\r\n", cnt, string(buf[:cnt]))
}
