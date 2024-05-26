package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()

	data := []byte("hello")
	//发送三次数据
	for i := 0; i < 3; i++ {
		//客户端发送数据
		conn.Write(data)

		//接受服务端的返回，但读取的数据可能超过buf大小，需要循环读取
		buf := make([]byte, 1) //512
		i := 0
		for {
			cnt, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("读取文件结尾")
					break
				}
			}
			i++
			fmt.Printf("读取次数:%d\r\n", i)
			fmt.Printf("服务端数据长度:%d, 数据:%s\r\n", cnt, string(buf[:cnt]))
		}
	}
}
