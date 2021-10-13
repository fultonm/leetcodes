package main

func minSumSubArray(nums []int) int {
	const IntMax = int(^uint(0) >> 1)
	minInt := func(args ...int) int {
		r := IntMax
		for _, e := range args {
			if e < r {
				r = e
			}
		}
		return r
	}
	n := len(nums)
	min := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = minInt(nums[i-1]+nums[i], nums[i])
		min = minInt(min, nums[i])
	}
	return min
}
