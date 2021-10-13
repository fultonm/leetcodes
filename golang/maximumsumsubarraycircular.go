package main

func maxSubarraySumCircular(nums []int) int {
	sumInt := func(args ...int) int {
		r := 0
		for _, e := range args {
			r += e
		}
		return r
	}
	const IntMin = -int(^uint(0)>>1) - 1
	maxInt := func(args ...int) int {
		r := IntMin
		for _, e := range args {
			if e > r {
				r = e
			}
		}
		return r
	}
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
	maxSubArray := func(nums []int) int {
		const IntMin = -int(^uint(0)>>1) - 1
		maxInt := func(args ...int) int {
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
			nums[i] = maxInt(nums[i-1]+nums[i], nums[i])
			max = maxInt(max, nums[i])
		}
		return max
	}
	max := maxSubArray(append([]int{}, nums...))
	if max < 0 {
		return max
	}
	sum := sumInt(nums...)
	min := minInt(0, minSumSubArray(nums))
	return maxInt(max, sum-min)
}
