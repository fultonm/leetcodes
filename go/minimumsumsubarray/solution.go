package minimumsumsubarray

import "main/helper"

func MinSumSubArray(nums []int) int {
	return minSumSubArray(nums)
}

func minSumSubArray(nums []int) int {
	n := len(nums)
	min := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = helper.MinInt(nums[i-1]+nums[i], nums[i])
		min = helper.MinInt(min, nums[i])
	}
	return min
}
