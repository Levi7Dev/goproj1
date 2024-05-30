package main

import (
	"fmt"
	"net/http"
)

func main() {
	//先注册路由
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)

		_, err := w.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err)
		}
	})

	//再启动http服务
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http start err:", err)
		return
	}
}
