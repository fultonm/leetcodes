package main

import (
	"testing"
)

func TestMaximumSumSubarrayCircular(t *testing.T) {
	input := []int{4, -3, 2, 2}
	AssertInt(maxSubarraySumCircular(input), 8, t)

	input = []int{4, 1, 2, 2}
	AssertInt(maxSubarraySumCircular(input), 9, t)

	input = []int{5, 5, 0, -5, 3, -3, 2}
	AssertInt(maxSubarraySumCircular(input), 12, t)

	input = []int{2, -2, 2, 7, 8, 0}
	AssertInt(maxSubarraySumCircular(input), 19, t)

	input = []int{1, -2, 3, -2}
	AssertInt(maxSubarraySumCircular(input), 3, t)

	input = []int{-2, -3, -1}
	AssertInt(maxSubarraySumCircular(input), -1, t)
}