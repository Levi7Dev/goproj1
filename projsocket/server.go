package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := "127.0.0.1"
	port := "8848"
	address := fmt.Sprintf("%s:%s", ip, port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	fmt.Println("开启监听...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}

	//存储客户端发送的数据，字节切片
	buf := make([]byte, 1024)
	//cnt是实际读取到的数据
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}

	//打印数据
	fmt.Printf("数据长度：%d, 数据内容：%s\r\n", cnt, string(buf))

	//将数据转大写
	upperData := strings.ToUpper(string(buf[:cnt]))

	//返回客户端
	conn.Write([]byte(upperData))

	//关闭连接
	conn.Close()
}
