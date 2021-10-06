package main

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = MaxInt(nums[i-1]+nums[i], nums[i])
		max = MaxInt(max, nums[i])
	}
	return max
}
