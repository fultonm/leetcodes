package main

import "fmt"

func main() {
	test := []int{1, 2, 3}
	fmt.Printf("%v", getConcatenation(test))
}

func getConcatenation(nums []int) []int {
	x := len(nums)
	r := make([]int, x*2)
	for i := 0; i < x; i++ {
		r[i] = nums[i]
		r[i+x] = nums[i]
	}
	return r
}
