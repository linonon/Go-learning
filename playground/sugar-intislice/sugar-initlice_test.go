package sugar

import (
	"fmt"
	"testing"
)

func TestInitSliceWithSugar(t *testing.T) {
	var s = [...]int{0: -1, 10: -1}
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	ts := s[:]
	fmt.Printf("%T\n", ts)
	fmt.Printf("%v\n", ts)
}
