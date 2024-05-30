package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}

	rep, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	content := rep.Header.Get("Content-Type")
	date := rep.Header.Get("Date")

	fmt.Println(content)
	fmt.Println(date)

	//获得io
	body := rep.Body
	readBytes, _ := io.ReadAll(body)
	fmt.Println("bytes:", string(readBytes)) //html

}
