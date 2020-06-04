package zrpc

import (
	"github.com/rooobot/zrpc/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterService(t *testing.T) {
	s := NewServer()
	err := s.RegisterService("helloworld", new(testdata.Service))
	assert.Nil(t, err)
}
