package stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerStream(t *testing.T) {

}

func TestWithMethod(t *testing.T) {
	var ss ServerStream
	ss.WithMethod("test")
	assert.Equal(t, "test", ss.Method)
}

func TestClone(t *testing.T) {
	var ss ServerStream
	ss.Method = "test"
	test := ss.Clone()
	assert.Equal(t, ss.Method, test.Method)
}
