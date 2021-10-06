package main

import (
	"testing"
)

func TestPrisonCells(t *testing.T) {
	input := []int{0, 1, 0, 1, 1, 0, 0, 1}
	AssertIntArr(prisonAfterNDays(input, 7), []int{0, 0, 1, 1, 0, 0, 0, 0}, t)
}
