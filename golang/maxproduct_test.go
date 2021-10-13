package main

import (
	"golang/helper"
	"testing"
)

func TestMaxProductSubarray(t *testing.T) {
	input := []int{2, 3, -2, 4}
	helper.AssertInt(maxProduct(input), 6, t)

	input = []int{-2, 0, -1}
	helper.AssertInt(maxProduct(input), 0, t)

	input = []int{-3, -1, -1}
	helper.AssertInt(maxProduct(input), 3, t)
}
