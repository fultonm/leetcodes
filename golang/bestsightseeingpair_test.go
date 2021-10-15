package main

import (
	"golang/helper"
	"testing"
)

func TestBestSightseeingPair(t *testing.T) {
	input := []int{8, 1, 5, 6, 2}
	helper.AssertInt(maxScoreSightseeingPair(input), 11, t)

	input = []int{2, 7, 7, 2, 1, 7, 10, 4, 3, 3}
	helper.AssertInt(maxScoreSightseeingPair(input), 16, t)

}
