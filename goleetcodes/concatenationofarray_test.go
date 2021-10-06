package main

import (
	"testing"
)

func TestConcatArr(t *testing.T) {
	input := []int{1, 2, 3}
	AssertIntArr(getConcatenation(input), []int{1, 2, 3, 1, 2, 3}, t)

	input = []int{1, 3, 2, 1}
	AssertIntArr(getConcatenation(input), []int{1, 3, 2, 1, 1, 3, 2, 1}, t)

	input = []int{}
	AssertIntArr(getConcatenation(input), []int{}, t)
}
