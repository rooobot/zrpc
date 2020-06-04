package main

import (
	"fmt"
	"github.com/rooobot/zrpc"
	zhttp "github.com/rooobot/zrpc/http"
	"net/http"
	"strings"
	"time"
)

func init() {
	zhttp.HandleFunc("GET", "/hello", sayHello)
}

func main() {
	opts := []zrpc.ServerOption{
		zrpc.WithAddress("127.0.0.1:8000"),
		zrpc.WithProtocol("http"),
		zrpc.WithNetwork("tcp"),
		zrpc.WithTimeout(time.Millisecond * 2000),
	}
	s := zrpc.NewServer(opts ...)
	s.ServeHttp()
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	w.Write([]byte("world"))
}
