package stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientWithMethod(t *testing.T) {
	var cs ClientStream
	cs.WithMethod("test")
	assert.Equal(t, "test", cs.Method)
}

func TestClientClone(t *testing.T) {
	var cs ClientStream
	cs.Method = "test"
	test := cs.Clone()
	assert.Equal(t, cs.Method, test.Method)
}

func TestClientWithServicename(t *testing.T) {
	var cs ClientStream
	cs.WithServiceName("test")
	assert.Equal(t, "test", cs.ServiceName)
}
