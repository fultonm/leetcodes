package main

func maxSubarraySumCircular(nums []int) int {
	max := maxSubArray(append([]int{}, nums...))
	if max < 0 {
		return max
	}
	sum := SumInt(nums...)
	min := MinInt(0, minSumSubArray(nums))
	return MaxInt(max, sum-min)
}
