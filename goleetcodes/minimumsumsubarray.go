package main

func minSumSubArray(nums []int) int {
	n := len(nums)
	min := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = MinInt(nums[i-1]+nums[i], nums[i])
		min = MinInt(min, nums[i])
	}
	return min
}
