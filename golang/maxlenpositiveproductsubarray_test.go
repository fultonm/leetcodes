package main

import (
	"golang/helper"
	"testing"
)

func TestMaxProductSubarrayLen(t *testing.T) {
	input := []int{1, -2, -3, 4}
	helper.AssertInt(getMaxLen(input), 4, t)

	input = []int{0, 1, -2, -3, -4}
	helper.AssertInt(getMaxLen(input), 3, t)

	input = []int{-1, -2, -3, 0, 1}
	helper.AssertInt(getMaxLen(input), 2, t)
}
