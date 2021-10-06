package main

import (
	"testing"
)

func TestMatrix01(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	AssertIntArr2d(updateMatrix(input), [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}, t)
}
