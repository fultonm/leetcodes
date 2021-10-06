package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := []int{-2, 1, 2, 0, 2, -4, 2}
	AssertInt(minSumSubArray(input), -4, t)
}
