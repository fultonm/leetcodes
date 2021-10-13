package helper

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

func AssertIntArr2dUnordered(result [][]int, expect [][]int, t *testing.T) {
	hR, hE := len(result), len(expect)
	if hR != hE {
		t.Fatalf("Expected height %v does not match result height %v", hR, hE)
	}
	for i := 0; i < hR; i++ {
		wR, wE := len(result[i]), len(expect[i])
		if wR != wE {
			t.Fatalf("Expected width %v does not match result width %v", hR, hE)
		}
		if !IntArr2DContainsIntArr(result, expect[i]) {
			t.Fatalf("Expected result: %v not contained in results", expect[i])
		}
	}
}

func AssertStringArr(result []string, expect []string, t *testing.T) {
	if len(result) != len(expect) {
		t.Fatalf("Result length %v does not equal expected length %v", len(result), len(expect))
	}
	for i := 0; i < len(result); i++ {
		if !StringArrContainsString(expect, result[i]) {
			t.Fatalf("Result %v not contained in expected values", result[i])
		}
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

func IntArr2DContainsIntArr(arr2d [][]int, arr []int) bool {
OUTER:
	for i := 0; i < len(arr2d); i++ {
		if len(arr) != len(arr2d[i]) {
			continue
		}
		for j := 0; j < len(arr); j++ {
			if arr[j] != arr2d[i][j] {
				continue OUTER
			}
		}
		return true
	}
	return false
}

func StringArrContainsString(arr []string, s string) bool {
	for _, e := range arr {
		if e == s {
			return true
		}
	}
	return false
}

func SumInt(args ...int) int {
	r := 0
	for _, e := range args {
		r += e
	}
	return r
}

const IntMax = int(^uint(0) >> 1)
const IntMin = -int(^uint(0)>>1) - 1

func MaxInt(args ...int) int {
	if len(args) == 0 {
		return IntMax
	}
	r := IntMin
	for _, e := range args {
		if e > r {
			r = e
		}
	}
	return r
}

func MinInt(args ...int) int {
	if len(args) == 0 {
		return IntMin
	}
	r := IntMax
	for _, e := range args {
		if e < r {
			r = e
		}
	}
	return r
}

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
