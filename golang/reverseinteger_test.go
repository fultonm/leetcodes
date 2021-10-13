package main

import (
	"golang/helper"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	helper.AssertInt(reverse(123), 321, t)
}
