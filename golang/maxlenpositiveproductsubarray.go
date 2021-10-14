package main

type Marker struct {
	left int
	len  int
}

func getMaxLen(nums []int) int {
	return mlps(nums, 0, len(nums)).len
}

func mlps(nums []int, j int, k int) Marker {
	/** Helpers */
	const IntMax = int(^uint(0) >> 1)
	const IntMin = -int(^uint(0)>>1) - 1
	MaxInt := func(args ...int) int {
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
	MinInt := func(args ...int) int {
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
	MaxLenMarker := func(args ...Marker) Marker {
		r := Marker{-1, 0}
		if len(args) == 0 {
			return r
		}
		for _, e := range args {
			if e.len > r.len {
				r = e
			}
		}
		return r
	}
	MarkersValid := func(args ...Marker) bool {
		for _, e := range args {
			if e.left < 0 || e.len < 0 {
				return false
			}
		}
		return true
	}
	/** Solution */
	n := k - j
	if n == 0 {
		return Marker{left: -1, len: 0}
	} else if n == 1 {
		if nums[j] > 0 {
			return Marker{left: j, len: 1}
		} else {
			return Marker{left: -1, len: 0}
		}
	}
	m := j + ((k - j) / 2)
	maxLeft := mlps(nums, j, m)
	maxRight := mlps(nums, m, k)
	// cross left
	leftNegCount, leftmostNeg, leftmostPos := 0, IntMax, IntMax
	for i := m - 1; i >= j; i-- {
		if nums[i] == 0 {
			break
		} else if nums[i] < 0 {
			leftNegCount++
			leftmostNeg = i
		} else {
			leftmostPos = i
		}
	}
	leftmostPosProdIdx := IntMax
	leftmostNegProdIdx := IntMax
	if leftNegCount%2 == 0 {
		leftmostPosProdIdx = MinInt(leftmostNeg, leftmostPos)
		if leftNegCount > 0 {
			leftmostNegProdIdx = leftmostNeg + 1
		}
	} else {
		leftmostNegProdIdx = MinInt(leftmostNeg, leftmostPos)
		if leftmostNegProdIdx < m-1 {
			leftmostPosProdIdx = leftmostNeg + 1
		}
	}
	maxCrossLeftPosProd := Marker{left: -1, len: 0}
	if leftmostPosProdIdx != IntMax {
		maxCrossLeftPosProd = Marker{
			left: leftmostPosProdIdx,
			len:  m - leftmostPosProdIdx,
		}
	}
	maxCrossLeftNegProd := Marker{left: -1, len: 0}
	if leftmostNegProdIdx != IntMax {
		maxCrossLeftNegProd = Marker{
			left: leftmostNegProdIdx,
			len:  m - leftmostNegProdIdx,
		}
	}
	// cross right
	rightNegCount, rightmostNeg, rightmostPos := 0, IntMin, IntMin
	for i := m; i < k; i++ {
		if nums[i] == 0 {
			break
		} else if nums[i] < 0 {
			rightNegCount++
			rightmostNeg = i
		} else {
			rightmostPos = i
		}
	}
	rightmostPosProdIdx := IntMin
	rightmostNegProdIdx := IntMin
	if rightNegCount%2 == 0 {
		rightmostPosProdIdx = MaxInt(rightmostNeg, rightmostPos)
		if rightNegCount > 0 {
			rightmostNegProdIdx = rightmostNeg - 1
		}
	} else {
		rightmostNegProdIdx = MaxInt(rightmostNeg, rightmostPos)
		if rightmostNegProdIdx > m {
			rightmostPosProdIdx = rightmostNeg - 1
		}
	}
	maxCrossRightPosProd := Marker{left: -1, len: 0}
	if rightmostPosProdIdx != IntMin {
		maxCrossRightPosProd = Marker{
			left: m,
			len:  rightmostPosProdIdx - m + 1,
		}
	}
	maxCrossRightNegProd := Marker{-1, -1}
	if rightmostNegProdIdx != IntMin {
		maxCrossRightNegProd = Marker{
			left: m,
			len:  rightmostNegProdIdx - m + 1,
		}
	}
	maxCross := Marker{left: -1, len: 0}
	if MarkersValid(maxCrossLeftPosProd, maxCrossRightPosProd, maxCrossLeftNegProd, maxCrossRightNegProd) {
		if maxCrossLeftPosProd.len+maxCrossRightPosProd.len > maxCrossLeftNegProd.len+maxCrossRightNegProd.len {
			maxCross = Marker{
				left: maxCrossLeftPosProd.left,
				len:  maxCrossLeftPosProd.len + maxCrossRightPosProd.len,
			}
		} else {
			maxCross = Marker{
				left: maxCrossLeftNegProd.left,
				len:  maxCrossLeftNegProd.len + maxCrossRightNegProd.len,
			}
		}
	} else if MarkersValid(maxCrossLeftPosProd, maxCrossRightPosProd) {
		maxCross = Marker{
			left: maxCrossLeftPosProd.left,
			len:  maxCrossLeftPosProd.len + maxCrossRightPosProd.len,
		}
	} else if MarkersValid(maxCrossLeftNegProd, maxCrossRightNegProd) {
		maxCross = Marker{
			left: maxCrossLeftNegProd.left,
			len:  maxCrossLeftNegProd.len + maxCrossRightNegProd.len,
		}
	}
	return MaxLenMarker(maxLeft, maxRight, maxCross)
}
