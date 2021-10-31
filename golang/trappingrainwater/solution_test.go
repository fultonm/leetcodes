package trappingrainwater

import (
	"golang/helper"
	"testing"
)

func Test(t *testing.T) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	helper.AssertInt(trap(input), 6, t)
}
