package codewar

import (
	"fmt"
	"strings"
	"testing"
)

func stringEndsWith(str, ending string) bool {
	return ending == "" || strings.Contains(str, ending) && strings.LastIndex(str, ending)+len(ending) == len(str)
}

func stringEndsWithBP(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func Test(t *testing.T) {
	fmt.Println(stringEndsWith("", ""))
	fmt.Println(stringEndsWith(" ", ""))
	fmt.Println(stringEndsWith("abc", "c"))
	fmt.Println(stringEndsWith("sensei", "i"))
	fmt.Println(stringEndsWith("$a = $b + 1", "+1"))
	fmt.Println(stringEndsWith("sensei", "se"))

	fmt.Println(stringEndsWithBP("", ""))
	fmt.Println(stringEndsWithBP(" ", ""))
	fmt.Println(stringEndsWithBP("abc", "c"))
	fmt.Println(stringEndsWithBP("sensei", "i"))
	fmt.Println(stringEndsWithBP("$a = $b + 1", "+1"))
	fmt.Println(stringEndsWithBP("sensei", "se"))
}
