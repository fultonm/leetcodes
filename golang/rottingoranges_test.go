package main

import "testing"

func TestRottingOranges(t *testing.T) {
	input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	AssertInt(orangesRotting(input), 4, t)

	input = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	AssertInt(orangesRotting(input), -1, t)

	input = [][]int{{0, 2}}
	AssertInt(orangesRotting(input), 0, t)

}
