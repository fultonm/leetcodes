package main

import (
	"math/rand"
	"testing"
	"time"
)

func AssertInt(result int, expect int, t *testing.T) {
	if result != expect {
		t.Fatalf("Expected %v, got %v", expect, result)
	}
}

func AssertIntArr(result []int, expect []int, t *testing.T) {
	wR, wE := len(result), len(expect)
	if wR != wE {
		t.Fatalf("Expected width %v does not match result width %v", wR, wE)
	}
	for i := 0; i < wR; i++ {
		AssertInt(result[i], expect[i], t)
	}
}

func AssertIntArr2d(result [][]int, expect [][]int, t *testing.T) {
	hR, hE := len(result), len(expect)
	if hR != hE {
		t.Fatalf("Expected height %v does not match result height %v", hR, hE)
	}
	for i := 0; i < hR; i++ {
		AssertIntArr(result[i], expect[i], t)
	}
}

func RandomArray(length int, min int, max int) []int {
	arr := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr[i] = min + rand.Intn(max-min+1)
	}
	return arr
}

func StringContainsRune(s string, r rune) int {
	for i, e := range []rune(s) {
		if e == r {
			return i
		}
	}
	return -1
}

func RuneArrContainsRune(arr []rune, r rune) int {
	for i, e := range arr {
		if e == r {
			return i
		}
	}
	return -1
}

func SumInt(args ...int) int {
	r := 0
	for _, e := range args {
		r += e
	}
	return r
}

func MaxInt(args ...int) int {
	r := IntMin
	for _, e := range args {
		if e > r {
			r = e
		}
	}
	return r
}

const IntMin = -int(^uint(0)>>1) - 1

func MinInt(args ...int) int {
	r := IntMax
	for _, e := range args {
		if e < r {
			r = e
		}
	}
	return r
}

const IntMax = int(^uint(0) >> 1)

func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func SelectionSortInt(arr []int) {
	findMinIdx := func(s int, arr []int) int {
		min := arr[s]
		idx := -1
		for i, e := range arr {
			if e < min {
				min = e
				idx = i
			}
		}
		return idx
	}

	size := len(arr)
	for i := 0; i < size; i++ {
		minIdx := findMinIdx(i, arr)
		if minIdx > -1 {
			tmp := arr[i]
			arr[i] = arr[minIdx]
			arr[minIdx] = tmp
		}
	}
}
