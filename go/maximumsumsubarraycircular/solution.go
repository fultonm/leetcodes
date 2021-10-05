package maximumsumsubarraycircular

import (
	"main/helper"
	"main/maximumsubarray"
	"main/minimumsumsubarray"
)

func maxSubarraySumCircular(nums []int) int {
	max := maximumsubarray.MaxSubArray(append([]int{}, nums...))
	if max < 0 {
		return max
	}
	sum := helper.SumInt(nums...)
	min := helper.MinInt(0, minimumsumsubarray.MinSumSubArray(nums))
	return helper.MaxInt(max, sum-min)
}
