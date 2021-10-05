package maximumsubarray

import "main/helper"

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = helper.MaxInt(nums[i-1]+nums[i], nums[i])
		max = helper.MaxInt(max, nums[i])
	}
	return max
}
