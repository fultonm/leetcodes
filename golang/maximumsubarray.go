package main

func maxSubArray(nums []int) int {
	const IntMin = -int(^uint(0)>>1) - 1
	MaxInt := func(args ...int) int {
		r := IntMin
		for _, e := range args {
			if e > r {
				r = e
			}
		}
		return r
	}
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = MaxInt(nums[i-1]+nums[i], nums[i])
		max = MaxInt(max, nums[i])
	}
	return max
}
