package main

import (
	"context"
	"fmt"
	"github.com/rooobot/zrpc/client"
	"github.com/rooobot/zrpc/testdata"
	"time"
)

func main() {
	opts := []client.Option{
		client.WithTarget("127.0.0.1:8000"),
		client.WithNetwork("tcp"),
		client.WithSerializationType("msgpack"),
		client.WithTimeout(time.Millisecond * 2000),
	}
	c := client.DefaultClient
	req := &testdata.HelloRequest{Msg: "hello"}
	rsp := &testdata.HelloReply{}
	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts...)
	fmt.Println(rsp.Msg, err)
}
