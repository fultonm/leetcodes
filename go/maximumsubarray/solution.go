package maximumsubarray

import "main/helper"

func MSA(nums []int) int {
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	n := len(nums)
	prev := nums[0]
	max := prev
	for i := 1; i < n; i++ {
		prev = helper.MaxInt(prev+nums[i], nums[i])
		if prev > max {
			max = prev
		}
	}
	return max
}
