package main

import (
	"context"
	"fmt"
	"github.com/rooobot/zrpc"
	"github.com/rooobot/zrpc/examples/helloworld2/helloworld"
	"time"
)

type greeterService struct{}

func (g *greeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("recv Msg: ", req.Msg)
	rsp := &helloworld.HelloReply{
		Msg: req.Msg + " world",
	}
	return rsp, nil
}

func main() {
	opts := []zrpc.ServerOption{
		zrpc.WithAddress("127.0.0.1:8000"),
		zrpc.WithNetwork("tcp"),
		zrpc.WithProtocol("proto"),
		zrpc.WithTimeout(2000 * time.Millisecond),
	}
	s := zrpc.NewServer(opts...)
	helloworld.
}