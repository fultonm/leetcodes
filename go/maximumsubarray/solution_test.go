package maximumsubarray

import (
	"main/helper"
	"testing"
)

func Test(t *testing.T) {
	input := []int{-2, 1, 2, 0, 2, -4, 2}
	result, expect := maxSubArray(input), 5
	helper.Assert(result, expect, t)

	input = []int{1}
	helper.Assert(maxSubArray(input), 1, t)

	input = []int{5, 5, 0, -5, 3, -3, 2}
	helper.Assert(maxSubArray(input), 10, t)
}
