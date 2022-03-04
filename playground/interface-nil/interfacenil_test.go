package playground

import (
	"fmt"
	"testing"
)

type People interface {
	Show() string
}

type Student struct{}

func (s *Student) Show() string {
	return ""
}

type PeopelWithoutMethod interface{}

type Student2 struct {
	PeopelWithoutMethod
}

func live() People {
	var stu *Student
	return stu
}

func live2() PeopelWithoutMethod {
	var stu *Student2
	return stu
}

func TestInterfaceNil(t *testing.T) {
	if live2() == nil {
		fmt.Println("AAAAAAAAA")
	} else {
		fmt.Println("BBBBBBBBB")
	}
}
