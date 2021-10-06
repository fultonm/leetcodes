package main

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	AssertInt(reverse(123), 321, t)
}
