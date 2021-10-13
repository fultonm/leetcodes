package main

func maxProduct(nums []int) int {
	/** Helpers */
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
	/** Solution start */
	n := len(nums)
	if n == 0 {
		return IntMin
	} else if n == 1 {
		return nums[0]
	}
	m := n / 2
	maxLeft := maxProduct(nums[:m])
	maxRight := maxProduct(nums[m:])
	productCrossLeft := nums[m-1]
	maxCrossLeft, minCrossLeft := productCrossLeft, productCrossLeft
	for i := m - 2; i >= 0; i-- {
		productCrossLeft *= nums[i]
		maxCrossLeft = MaxInt(maxCrossLeft, productCrossLeft)
		minCrossLeft = MinInt(minCrossLeft, productCrossLeft)
	}
	productCrossRight := nums[m]
	maxCrossRight, minCrossRight := productCrossRight, productCrossRight
	for i := m + 1; i < n; i++ {
		productCrossRight *= nums[i]
		maxCrossRight = MaxInt(maxCrossRight, productCrossRight)
		minCrossRight = MinInt(minCrossRight, productCrossRight)
	}
	maxCross := MaxInt(minCrossLeft*minCrossRight,
		maxCrossLeft*maxCrossRight)
	return MaxInt(maxLeft, maxRight, maxCross)
}
