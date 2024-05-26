package main

import (
	"fmt"
	"io"
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
	defer listener.Close()

	fmt.Println("开启监听...")

	//循环监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			return
		}

		go handle(conn)
	}
}

// 处理某个tcp连接
func handle(conn net.Conn) {
	defer conn.Close()
	//存储客户端发送的数据，字节切片
	buf := make([]byte, 1024)

	//一个连接可能发送多条数据，buf可以复用，当读取数据超过buf的大小会在下一次循环中反复读取
	for {
		//cnt是实际读取到的数据
		cnt, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconected...")
				break
			}
			fmt.Println("read err:", err)
			return
		}

		//打印数据
		fmt.Printf("数据长度：%d, 数据内容：%s\r\n", cnt, string(buf[:cnt]))

		//将数据转大写
		upperData := strings.ToUpper(string(buf[:cnt]))

		//返回客户端
		conn.Write([]byte(upperData))
	}
}
