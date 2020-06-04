package main

import (
	"context"
	"fmt"
	"github.com/rooobot/zrpc/client"
	hw "github.com/rooobot/zrpc/examples/helloworld2/helloworld"
	"time"
)

func main() {
	opts := []client.Option{
		client.WithTarget("127.0.0.1:8000"),
		client.WithNetwork("tcp"),
		client.WithTimeout(2000 * time.Millisecond),
	}
	proxy := hw.NewGreeterClientProxy(opts...)
	req := &hw.HelloRequest{
		Msg: "hello",
	}
	rsp, err := proxy.SayHello(context.Background(), req, opts...)
	fmt.Println(rsp, err)
}
