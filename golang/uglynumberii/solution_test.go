package uglynumberii

import (
	"golang/helper"
	"testing"
)

func Test(t *testing.T) {
	helper.AssertInt(nthUglyNumber(10), 12, t)
	helper.AssertInt(nthUglyNumber(1), 1, t)
	helper.AssertInt(nthUglyNumber(1690), 1, t)
}
