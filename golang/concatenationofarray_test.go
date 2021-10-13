package main

import (
	"golang/helper"
	"testing"
)

func TestConcatArr(t *testing.T) {
	input := []int{1, 2, 3}
	helper.AssertIntArr(getConcatenation(input), []int{1, 2, 3, 1, 2, 3}, t)

	input = []int{1, 3, 2, 1}
	helper.AssertIntArr(getConcatenation(input), []int{1, 3, 2, 1, 1, 3, 2, 1}, t)

	input = []int{}
	helper.AssertIntArr(getConcatenation(input), []int{}, t)
}
