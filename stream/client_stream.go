package stream

import "context"

const ClientStreamKey = StreamContextKey("ZRPC_CLIENT_STREAM")

type ClientStream struct {
	ctx         context.Context
	ServiceName string // 服务名称
	Method      string // 方法名
}

func GetClientStream(ctx context.Context) *ClientStream {
	v := ctx.Value(ClientStreamKey)
	if v == nil {
		cs := &ClientStream{
			ctx: ctx,
		}
		context.WithValue(ctx, ClientStreamKey, cs)
	}
	return v.(*ClientStream)
}

func (cs *ClientStream) Clone() *ClientStream {
	return &ClientStream{
		Method: cs.Method,
	}
}

func NewClientStream(ctx context.Context) (context.Context, *ClientStream) {
	var cs *ClientStream
	v := ctx.Value(ClientStreamKey)
	if v != nil {
		cs = v.(*ClientStream)
	} else {
		cs = &ClientStream{
			ctx: ctx,
		}
	}
	valueCtx := context.WithValue(ctx, ClientStreamKey, cs)
	return valueCtx, cs
}

func (cs *ClientStream) WithMethod(method string) {
	cs.Method = method
}

func (cs *ClientStream) WithServiceName(serviceName string) {
	cs.ServiceName = serviceName
}
