package main

import (
	"testing"
)

func TestStringContainsSubstring(t *testing.T) {
	test1 := "abc"
	char1 := rune('s')
	if StringContainsRune(test1, char1) > 0 {
		t.Fatalf("Did not expected string: %v to have element %v\n", test1, string(char1))
	}
	char2 := rune('c')
	if StringContainsRune(test1, char2) < 0 {
		t.Fatalf("Expected string: %v to have element %v\n", test1, string(char2))

	}
}

func TestRuneArrContainsRune(t *testing.T) {
	test1 := []rune("abc")
	if RuneArrContainsRune(test1, rune('c')) < 0 {
		t.Fatalf("Expected rune array: %v to have element %v\n", test1, rune('c'))
	}
	if RuneArrContainsRune(test1, rune('d')) > 0 {
		t.Fatalf("Did not expect rune array: %v to have element %v\n", test1, rune('d'))
	}
}
