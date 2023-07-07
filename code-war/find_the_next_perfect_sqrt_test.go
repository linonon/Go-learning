package codewar

import (
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {

	num := 81
	s := math.Sqrt(float64(num))
	if int(s)*int(s) != num {
		fmt.Println(-1)
	}

	fmt.Println(s)
}

func TestBP(t *testing.T) {
	var num int64 = 81
	res := math.Pow(math.Sqrt(float64(num))+1, 2)
	if res != math.Trunc(res) {
		fmt.Println(-1)
	}
	fmt.Println(int64(res))
}
