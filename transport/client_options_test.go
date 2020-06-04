package transport

import (
	"github.com/rooobot/zrpc/pool/connpool"
	"github.com/rooobot/zrpc/selector"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWithServiceName(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithServiceName("test")
	fCto(&cto)
	assert.Equal(t, "test", cto.ServiceName)
	fCto = WithServiceName("")
	fCto(&cto)
	assert.Equal(t, "", cto.ServiceName)
}

func TestWithClientTarget(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithClientTarget("test")
	fCto(&cto)
	assert.Equal(t, "test", cto.Target)
	fCto = WithClientTarget("")
	fCto(&cto)
	assert.Equal(t, "", cto.Target)
}

func TestWithClientNetwork(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithClientNetwork("test")
	fCto(&cto)
	assert.Equal(t, "test", cto.Network)
	fCto = WithClientNetwork("")
	fCto(&cto)
	assert.Equal(t, "", cto.Network)
}

func TestWithClientPool(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithClientPool(connpool.GetPool("default"))
	fCto(&cto)
	assert.NotNil(t, cto.Pool)
}

func TestWithTimeout(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithTimeout(time.Second * time.Duration(2))
	fCto(&cto)
	assert.Equal(t, time.Second*time.Duration(2), cto.Timeout)
}

func TestWithSelector(t *testing.T) {
	var cto ClientTransportOptions
	fCto := WithSelector(selector.GetSelector("test"))
	fCto(&cto)
	assert.NotNil(t, cto.Selector)
}
