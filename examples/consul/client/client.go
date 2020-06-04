package main

import (
	"context"
	"fmt"
	"github.com/rooobot/zrpc/client"
	"github.com/rooobot/zrpc/plugin/consul"
	"github.com/rooobot/zrpc/testdata"
	"time"
)

func main() {
	opts := []client.Option {
		client.WithNetwork("tcp"),
		client.WithTimeout(2000 * time.Millisecond),
		client.WithSelectorName(consul.Name),
	}
	c := client.DefaultClient
	req := &testdata.HelloRequest{
		Msg: "hello",
	}
	rsp := &testdata.HelloReply{}

	consul.Init("localhost:8500")
	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts ...)
	fmt.Println(rsp.Msg, err)
}
