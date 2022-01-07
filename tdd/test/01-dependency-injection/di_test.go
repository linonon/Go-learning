package dependencyinjection_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Greet print name
func Greet(writer io.Writer, name string) {
	// In `Printf()`, there is a `hook in`
	fmt.Fprintf(writer, "Hello, %s", name)
}

// TestGreet is testing greet
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' what '%s'", got, want)
	}

}
