package client

import (
	"context"
	"github.com/rooobot/zrpc"
	"github.com/rooobot/zrpc/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCall(t *testing.T) {

	var ch = make(chan struct{})
	go func() {
		serverOpts := []zrpc.ServerOption{
			zrpc.WithAddress("127.0.0.1:8001"),
			zrpc.WithNetwork("tcp"),
			zrpc.WithSerializationType("msgpack"),
			zrpc.WithTimeout(time.Millisecond * 2000),
		}
		s := zrpc.NewServer(serverOpts ...)
		if err := s.RegisterService("helloworld.Greeter", new(testdata.Service)); err != nil {
			panic(err)
		}

		go func() {
			s.Serve()
		}()

		<- ch
		s.Close()
	}()

	time.Sleep(1000 * time.Millisecond)

	opts := []Option {
		WithTarget("127.0.0.1:8001"),
		WithNetwork("tcp"),
		WithTimeout(2000 * time.Millisecond),
		WithSerializationType("msgpack"),
	}
	c := DefaultClient
	req := &testdata.HelloRequest{
		Msg: "hello",
	}
	rsp := &testdata.HelloReply{}

	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts ...)

	close(ch)

	assert.Nil(t, err)
}
