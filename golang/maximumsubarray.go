package main

func maxSubArray(nums []int) int {
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
	n := len(nums)
	if n == 0 {
		return IntMin
	} else if n == 1 {
		return nums[0]
	}
	m := n / 2
	maxLeft := maxSubArray(nums[:m])
	maxRight := maxSubArray(nums[m:])
	sumCrossLeft := nums[m-1]
	maxCrossLeft := sumCrossLeft
	for i := m - 2; i >= 0; i-- {
		sumCrossLeft += nums[i]
		maxCrossLeft = MaxInt(maxCrossLeft, sumCrossLeft)
	}
	sumCrossRight := nums[m]
	maxCrossRight := sumCrossRight
	for i := m + 1; i < n; i++ {
		sumCrossRight += nums[i]
		maxCrossRight = MaxInt(maxCrossRight, sumCrossRight)
	}
	return MaxInt(maxLeft, maxRight, maxCrossLeft+maxCrossRight)
}
