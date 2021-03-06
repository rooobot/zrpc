package main

import (
	"github.com/rooobot/zrpc"
	"github.com/rooobot/zrpc/plugin/consul"
	"github.com/rooobot/zrpc/testdata"
	"time"
)

func main() {
	opts := []zrpc.ServerOption{
		zrpc.WithAddress("127.0.0.1:8000"),
		zrpc.WithNetwork("tcp"),
		zrpc.WithSerializationType("msgpack"),
		zrpc.WithTimeout(time.Millisecond * 2000),
		zrpc.WithSelectorSvrAddr("localhost:8500"),
		zrpc.WithPlugin(consul.Name),
	}
	s := zrpc.NewServer(opts ...)
	if err := s.RegisterService("helloworld.Greeter", new(testdata.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}
