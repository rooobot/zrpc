package main

import (
	"context"
	"errors"
	"github.com/rooobot/zrpc"
	"github.com/rooobot/zrpc/auth"
	"github.com/rooobot/zrpc/log"
	"github.com/rooobot/zrpc/metadata"
	"github.com/rooobot/zrpc/testdata"
	"time"
)

func main() {
	af := func(ctx context.Context) (context.Context, error) {
		md := metadata.ServerMetadata(ctx)

		if len(md) == 0 {
			return ctx, errors.New("token nil")
		}

		v := md["authorization"]
		log.Debug("token : " , string(v))
		if string(v) != "Bearer testToken" {
			return ctx, errors.New("token invalid")
		}
		return ctx, nil
	}

	opts := []zrpc.ServerOption{
		zrpc.WithAddress("127.0.0.1:8003"),
		zrpc.WithNetwork("tcp"),
		zrpc.WithSerializationType("msgpack"),
		zrpc.WithTimeout(time.Millisecond * 2000),
		zrpc.WithInterceptor(auth.BuildAuthInterceptor(af)),
	}
	s := zrpc.NewServer(opts...)
	if err := s.RegisterService("/helloworld.Greeter", new(testdata.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}