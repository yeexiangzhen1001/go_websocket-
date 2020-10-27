package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//创建路由
	router := mux.NewRouter()
	//ws控制器不断处理管道数据，进行同步数据
	go  h.run()
	//指定回调函数
	router.HandleFunc("/ws",wsHandler)
	//开启服务端监听
	if err := http.ListenAndServe("127.0.0.1:8080",router); err != nil {
		fmt.Println("err：",err)
	}
}
