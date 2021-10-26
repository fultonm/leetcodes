package besttimetobuyandsellstockwithtransactionfee

import (
	"golang/helper"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 3, 2, 8, 4, 9}
	helper.AssertInt(maxProfit(input, 2), 8, t)
}
