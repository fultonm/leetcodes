package main

import (
	"golang/helper"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	input := []int{-2, 1, 2, 0, 2, -4, 2}
	helper.AssertInt(maxSubArray(input), 5, t)

	input = []int{1}
	helper.AssertInt(maxSubArray(input), 1, t)

	input = []int{5, 5, 0, -5, 3, -3, 2}
	helper.AssertInt(maxSubArray(input), 10, t)

	input = []int{-2, -1, -1, -2, -7}
	helper.AssertInt(maxSubArray(input), -1, t)

	input = []int{1, -2, 3, -2}
	helper.AssertInt(maxSubArray(input), 3, t)

	input = []int{4, 1, 3, 5}
	helper.AssertInt(maxSubArray(input), 13, t)
}
