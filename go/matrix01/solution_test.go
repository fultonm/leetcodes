package matrix01

import (
	"main/helper"
	"testing"
)

func Test(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	helper.AssertIntArr2d(updateMatrix(input), [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}, t)
}
