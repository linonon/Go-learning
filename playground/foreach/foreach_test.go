package playground

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	stus := []student{
		{
			Name: "linonon",
			Age:  1,
		}, {
			Name: "lumwenon",
			Age:  2,
		}, {
			Name: "oliver",
			Age:  3,
		},
	}

	m := make(map[string]*student)
	for _, stu := range stus {
		fmt.Printf("%x", &stu)
		m[stu.Name] = &stu
	}

	fmt.Println("Foreach Wrong Answer")
	for k, v := range m {
		fmt.Println("\t", k, v)
	}

	m2 := make(map[string]*student)
	for i := range stus {
		m2[stus[i].Name] = &stus[i]
	}

	fmt.Println("Foreach Right Answer")
	for k, v := range m2 {
		fmt.Println("\t", k, v)
	}
}
