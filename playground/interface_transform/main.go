package main

import "fmt"

func printAny(args ...any) {
	a := args[0].(int)
	b := args[1].(float64)
	c := args[2].(string)
	fmt.Println(a, b, c)

	aErr, ok := args[0].(int8)
	if !ok {
		fmt.Println("aErr is not int8")
	}
	bErr, ok := args[1].(float32)
	if !ok {
		fmt.Println("bErr is not float32")
	}
	cErr, ok := args[2].(int)
	if !ok {
		fmt.Println("cErr is not int")
	}
	fmt.Println(aErr, bErr, cErr)

	// will panic
	aErr2 := args[0].(int8)
	_ = aErr2
}

func main() {
	var (
		a int     = 1
		b float64 = 2.0
		c string  = "3"
	)

	printAny(a, b, c)
}
