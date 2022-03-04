package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type IExample interface {
	Hello(n int) int
}

type Example struct {
}

func (e *Example) Hello(n int) int {
	fmt.Printf("Hello with %d\n", n)
	return n
}

// ExampleFunc for testify
func ExampleFunc(e IExample) {
	for n := 1; n <= 3; n++ {
		for i := 0; i < n; i++ {
			e.Hello(n)
		}
	}
}

type MockExample struct {
	mock.Mock
}

func (m *MockExample) Hello(n int) int {
	args := m.Called(n)
	return args.Int(0)
}

func TestExample(t *testing.T) {
	e := new(MockExample)

	e.On("Hello", 1).Return(1).Times(1)
	e.On("Hello", 2).Return(2).Times(2)
	e.On("Hello", 3).Return(3).Times(3)

	ExampleFunc(e)

	e.AssertExpectations(t)
}
