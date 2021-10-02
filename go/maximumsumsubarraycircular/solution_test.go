package maximumsumsubarraycircular

import (
	"main/helper"
	"testing"
)

func Test(t *testing.T) {
	// input := []int{4, -3, 2, 2}
	// helper.Assert(maxSubarraySumCircular(input), 8, t)

	// input = []int{4, 1, 2, 2}
	// helper.Assert(maxSubarraySumCircular(input), 9, t)

	input := []int{5, 5, 0, -5, 3, -3, 2}
	helper.Assert(maxSubarraySumCircular(input), 12, t)
}
