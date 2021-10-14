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

	input = []int{70, -18, 75, -72, -69, -84, 64, -65,
		0, -82, 62, 54, -63, -85, 53, -60, -59, 29, 32,
		59, -54, -29, -45, 0, -10, 22, 42, -37, -16, 0,
		-7, -76, -34, 37, -10, 2, -59, -24, 85, 45, -81,
		56, 86}
	helper.AssertInt(getMaxLen(input), 14, t)

	input = []int{6, 2, 10, 1, -2, 8}
	helper.AssertInt(getMaxLen(input), 4, t)

}
