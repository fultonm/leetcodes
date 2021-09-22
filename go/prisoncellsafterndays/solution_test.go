package prisoncellsafterndays

import (
	"reflect"
	"testing"
)

func throwIfNotEqual(expected []int, output []int, t *testing.T) {
	if !reflect.DeepEqual(expected, output) {
		t.Fatalf("Expected %v, got %v", expected, output)
	}
}

func Test1(t *testing.T) {
	input := []int{0, 1, 0, 1, 1, 0, 0, 1}
	n := 7
	expected := []int{0, 0, 1, 1, 0, 0, 0, 0}
	output := prisonAfterNDays(input, n)
	throwIfNotEqual(expected, output, t)
}
