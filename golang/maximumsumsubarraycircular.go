package main

func maxSubarraySumCircular(nums []int) int {
	const IntMax = int(^uint(0) >> 1)
	const IntMin = -int(^uint(0)>>1) - 1
	MaxInt := func(args ...int) int {
		if len(args) == 0 {
			return IntMax
		}
		r := IntMin
		for _, e := range args {
			if e > r {
				r = e
			}
		}
		return r
	}
	if len(nums) == 0 {
		return IntMax
	} else if len(nums) == 1 {
		return nums[0]
	}
	max := maxSubArray(nums)
	if max < 0 {
		return max
	}
	min := minSumSubArray(nums)
	sum := SumInt(nums...)
	return MaxInt(max, sum-min)
}
func SumInt(args ...int) int {
	r := 0
	for _, e := range args {
		r += e
	}
	return r
}
