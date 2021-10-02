package maximumsumsubarraycircular

import (
	"main/helper"
	"main/maximumsubarray"
)

// nums = [2, -3, 1]
// n = 3, n2 = 5, memo = [(1,2),(2, -1),(1, 1),,], max = (1,2), i = 2, sum = 0
func MSSC(nums []int) int {
	return maxSubarraySumCircular(nums)
}

// type Memo struct {
// 	len int
// 	sum int
// }

// func maxSubarraySumCircular(nums []int) int {
// 	n := len(nums)
// 	n2 := (n * 2) - 1
// 	memo := make([]Memo, n2)
// 	memo[0] = Memo{1, nums[0]}
// 	max := 0
// 	for i := 1; i < n2; i++ {
// 		// [4, -3, 2, 2][4, -3, 2, 2]
// 		//  0   1  2  3  4   5  6  7
// 		// i = 5, sum = 5, len = 4
// 		// sum = 5 - -3 = 8
// 		// len = 3
// 		// 8 > -3
// 		// memo[5] = 4, 8
// 		sum, len := memo[i-1].sum, memo[i-1].len
// 		if len >= n {
// 			sum -= nums[(i+1)%n]
// 			len--
// 		} else {
// 			sum += nums[i%n]
// 			len++
// 		}
// 		if sum > nums[i%n] {
// 			memo[i] = Memo{len, sum}
// 		} else {
// 			memo[i] = Memo{1, nums[i%n]}
// 		}
// 		if memo[i].sum > memo[max].sum {
// 			max = i
// 		}
// 	}
// 	return memo[max].sum
// }

func maxSubarraySumCircular(nums []int) int {
	mid := len(nums) / 2
	a := maximumsubarray.MSA(nums)
	b := maximumsubarray.MSA(append(nums[mid:], nums[:mid]...))
	return helper.MaxInt(a, b)
}
