package main

import (
	"Levi7Dev/goproj1/chatroom/common"
	"fmt"
	"net"
)

var allUsersMap = make(map[string]common.User)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	fmt.Println("建立监听成功...")

	//获取连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		//协程处理逻辑
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Println("处理逻辑...")

	//字节码缓存
	buf := make([]byte, 512)

	//多条消息，每条消息可能超过缓存大小，超出部分在下一个循环读取
	for {
		//获取远程连接ip和port作为id与name
		userAddr := conn.RemoteAddr().String()
		user := common.User{
			Id:   userAddr,
			Name: userAddr,
			Msg:  make(chan string),
		}
		//存入map中
		allUsersMap[user.Id] = user

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Printf("client===>server:读取到消息长度:%d\r\n", cnt)

		//TODO：消息发送逻辑
		fmt.Println(string(buf[:cnt]))
	}
}
