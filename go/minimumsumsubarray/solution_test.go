package minimumsumsubarray

import (
	"main/helper"
	"testing"
)

func Test(t *testing.T) {
	input := []int{-2, 1, 2, 0, 2, -4, 2}
	helper.AssertInt(minSumSubArray(input), -4, t)
}
