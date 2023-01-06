package main

import (
	"fmt"
	"log"
	"net/http"
)

// go build -gcflags "all=-N -l" game
// dlv --listen=:12345 --headless=true --api-version=2 --continue --accept-multiclient exec ./game &
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)
	
	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
