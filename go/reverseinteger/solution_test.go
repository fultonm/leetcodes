package reverseinteger

import (
	"testing"
)

func throwIfNotEqual(expected int, output int, t *testing.T) {
	if expected != output {
		t.Fatalf("Expected %d, got %d", expected, output)
	}
}

func Test1(t *testing.T) {
	input := 123
	expected := 321
	output := reverse(input)
	throwIfNotEqual(expected, output, t)
}
