package playground

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type Teacher struct{}

func (t *Teacher) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "u r a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestMethodSet(t *testing.T) {
	var peo People = &Teacher{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
