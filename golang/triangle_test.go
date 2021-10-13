package main

import (
	"golang/helper"
	"testing"
)

func TestTriangle(t *testing.T) {
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	output := 11
	helper.AssertInt(minimumTotal(input), output, t)
}
