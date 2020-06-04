package selector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var (
	mockConsulName = "mockConsul"
	nodeList       = []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}
)

type MockConsul struct {
	servicePool map[string][]string
}

func InitMockConsul() *MockConsul {
	return &MockConsul{
		servicePool: map[string][]string{
			"Greeter": nodeList,
		},
	}
}

func (c *MockConsul) Select(serviceName string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	nodeList, ok := c.servicePool[serviceName]
	if !ok {
		return "", fmt.Errorf("service not be registered in MockConsul")
	}

	index := rand.Int() % len(nodeList)
	return nodeList[index], nil
}

func TestGetSelector(t *testing.T) {
	m := InitMockConsul()
	RegisterSelector(mockConsulName, m)

	selector := GetSelector(mockConsulName)
	node, err := selector.Select("Greeter")
	assert.Nil(t, err)
	assert.Contains(t, nodeList, node)
	_, err = selector.Select("")
	assert.NotNil(t, err)

	selector = GetSelector("")
	assert.Equal(t, selector, DefaultSelector)
}
