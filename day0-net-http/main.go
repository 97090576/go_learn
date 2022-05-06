package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 将路由 / 的请求绑定到 handler
	http.HandleFunc("/", handler)
	// 将路由 /count 的请求绑定到 counter
	http.HandleFunc("/count", counter)
	// 启动服务，并监听 9999 端口
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}

var count = 0

func counter(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "you visit %q %d counts", r.URL.Path, count)
}
