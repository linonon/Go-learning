package main

import (
	"fmt"
	"strings"
)

var t1 = "This is an example!"
var w1 = "sihT si na !elpmaxe"
var t2 = "double  spaces"
var w2 = "elbuod  secaps"

func main() {
	tt1 := strings.Split(t1, " ")

	for i, str := range tt1 {
		rns := []rune(str)
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}

		tt1[i] = string(rns)
	}

	fmt.Println(strings.Join(tt1, " "))
}
