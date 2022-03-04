package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var expected = 100
	var actual = 100
	assert.Equal(t, expected, actual, "")
}

func TestEqualValue(t *testing.T) {
	type MyInt int
	var a = 100
	var b MyInt = 100
	// assert.Equal(t, a, b, "") // Fail
	assert.EqualValues(t, a, b, "") // Pass
}

