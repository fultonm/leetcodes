package main

import (
	"golang/helper"
	"testing"
)

func TestQuicksort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := append([]int{}, input...)
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)

	input = []int{10, 1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5, 10}
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)

	input = []int{-1}
	expected = []int{-1}
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)

	input = []int{}
	expected = []int{}
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)

	input = []int{1, 20, 20, 20}
	expected = []int{1, 20, 20, 20}
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)

	input = []int{1, 1, 1}
	expected = []int{1, 1, 1}
	Quicksort(&input)
	helper.AssertIntArr(input, expected, t)
}
