package concatenationofarray

import (
	"reflect"
	"testing"
)

func throwIfNotEqual(expected []int, concat []int, t *testing.T) {
	if !reflect.DeepEqual(expected, concat) {
		t.Fatalf("Expected %v, got %v", expected, concat)
	}
}

func Test1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 1, 2, 3}
	concat := getConcatenation(input)
	throwIfNotEqual(expected, concat, t)
}

func Test2(t *testing.T) {
	input := []int{1, 3, 2, 1}
	expected := []int{1, 3, 2, 1, 1, 3, 2, 1}
	concat := getConcatenation(input)
	throwIfNotEqual(expected, concat, t)
}

func TestEmpty(t *testing.T) {
	input := []int{}
	expected := []int{}
	concat := getConcatenation(input)
	throwIfNotEqual(expected, concat, t)
}
