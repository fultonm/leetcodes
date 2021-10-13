package main

func minSumSubArray(nums []int) int {
	const IntMax = int(^uint(0) >> 1)
	const IntMin = -int(^uint(0)>>1) - 1

	MinInt := func(args ...int) int {
		if len(args) == 0 {
			return IntMin
		}
		r := IntMax
		for _, e := range args {
			if e < r {
				r = e
			}
		}
		return r
	}
	n := len(nums)
	if n == 0 {
		return IntMax
	} else if n == 1 {
		return nums[0]
	}
	m := n / 2
	minLeft := minSumSubArray(nums[:m])
	minRight := minSumSubArray(nums[m:])
	sumCrossLeft := nums[m-1]
	minCrossLeft := sumCrossLeft
	for i := m - 2; i >= 0; i-- {
		sumCrossLeft += nums[i]
		minCrossLeft = MinInt(minCrossLeft, sumCrossLeft)
	}
	sumCrossRight := nums[m]
	minCrossRight := sumCrossRight
	for i := m + 1; i < n; i++ {
		sumCrossRight += nums[i]
		minCrossRight = MinInt(minCrossRight, sumCrossRight)
	}
	return MinInt(minLeft, minRight, minCrossLeft+minCrossRight)
}
