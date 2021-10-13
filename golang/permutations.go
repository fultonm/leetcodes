package main

func permute(nums []int) [][]int {
	r := [][]int{}
	n := len(nums)
	loopPermutations(0, n, &nums, &r)
	return r
}

func loopPermutations(l int, n int, nums *[]int, r *[][]int) {
	if l < n-1 {
		loopPermutations(l+1, n, nums, r)
		for i := l + 1; i < n; i++ {
			(*nums)[l], (*nums)[i] = (*nums)[i], (*nums)[l]
			loopPermutations(l+1, n, nums, r)
			(*nums)[l], (*nums)[i] = (*nums)[i], (*nums)[l]
		}
	} else {
		newPerm := make([]int, n)
		copy(newPerm, (*nums)[:])
		*r = append(*r, newPerm)
	}
}
