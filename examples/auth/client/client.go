package main

import (
	"context"
	"fmt"
	"github.com/rooobot/zrpc/auth"
	"github.com/rooobot/zrpc/client"
	"github.com/rooobot/zrpc/testdata"
	"time"
)

func main() {
	opts := []client.Option{
		client.WithTarget("127.0.0.1:8003"),
		client.WithNetwork("tcp"),
		client.WithTimeout(2000 * time.Millisecond),
		client.WithSerializationType("msgpack"),
		client.WithPerRPCAuth(auth.NewOAuth2ByToken("testToken")),
	}
	c := client.DefaultClient
	req := &testdata.HelloRequest{Msg: "hello"}
	rsp := &testdata.HelloReply{}
	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts...)
	fmt.Println(rsp.Msg, err)
}
