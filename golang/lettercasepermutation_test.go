package main

import (
	"golang/helper"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	input := "a1b2"
	expect := []string{"a1b2", "a1B2", "A1b2", "A1B2"}
	helper.AssertStringArr(letterCasePermutation(input), expect, t)

	input = "3z4"
	expect = []string{"3z4", "3Z4"}
	helper.AssertStringArr(letterCasePermutation(input), expect, t)

	input = "12345"
	expect = []string{"12345"}
	helper.AssertStringArr(letterCasePermutation(input), expect, t)

	input = "0"
	expect = []string{"0"}
	helper.AssertStringArr(letterCasePermutation(input), expect, t)
}
