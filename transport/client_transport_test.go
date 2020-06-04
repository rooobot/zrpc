package transport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var NewClientTransport = func() ClientTransport {
	return &clientTransport{
		opts: &ClientTransportOptions{ServiceName: "test"},
	}
}

func TestGetClientTransport(t *testing.T) {
	var testClientTransport = NewClientTransport()
	clientTransportMap["clinetTest"] = testClientTransport
	clientTransport := GetClientTransport("clinetTest")
	assert.Equal(t, clientTransport, testClientTransport)
	clientTransport = GetClientTransport("test")
	assert.Equal(t, clientTransport, DefaultClientTransport)
}
