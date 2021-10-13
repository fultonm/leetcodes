package main

import (
	"golang/helper"
	"testing"
)

func TestPermutations(t *testing.T) {
	input := []int{1, 2, 3}
	expect := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	helper.AssertIntArr2dUnordered(permute(input), expect, t)

	input = []int{0, 1}
	expect = [][]int{{0, 1}, {1, 0}}
	helper.AssertIntArr2dUnordered(permute(input), expect, t)

	input = []int{1}
	expect = [][]int{{1}}
	helper.AssertIntArr2dUnordered(permute(input), expect, t)
}
