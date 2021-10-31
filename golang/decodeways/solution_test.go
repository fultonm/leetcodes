package decodeways

import (
	"golang/helper"
	"testing"
)

func Test(t *testing.T) {
	input := "226"
	helper.AssertInt(numDecodings(input), 3, t)

	input = "10"
	helper.AssertInt(numDecodings(input), 1, t)

	input = "01"
	helper.AssertInt(numDecodings(input), 0, t)

	input = "12"
	helper.AssertInt(numDecodings(input), 2, t)

	input = "2101"
	helper.AssertInt(numDecodings(input), 1, t)
}
