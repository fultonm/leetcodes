package main

type SubarrayMarkers struct {
	val   int
	left  int
	right int
}

func getMaxLen(nums []int) int {
	subarray := maxLenPosProductSubarray(nums, 0, len(nums))
	return subarray.right - subarray.left
}

func maxLenPosProductSubarray(nums []int, left int, right int) SubarrayMarkers {
	/** Helpers */
	const IntMax = int(^uint(0) >> 1)
	const IntMin = -int(^uint(0)>>1) - 1
	MaxLenPosSubarrayMarkers := func(args ...SubarrayMarkers) SubarrayMarkers {
		if len(args) == 0 {
			return SubarrayMarkers{0, -1, -1}
		}
		r := SubarrayMarkers{0, -1, -1}
		for _, e := range args {
			if e.val > 0 && e.right-e.left > r.right-r.left {
				r = e
			}
		}
		return r
	}
	/** Solution start */
	n := right - left
	if n == 0 {
		return SubarrayMarkers{0, -1, -1}
	} else if n == 1 {
		if nums[left] > 0 {
			return SubarrayMarkers{nums[left], left, right}
		} else {
			return SubarrayMarkers{0, -1, -1}
		}
	}
	m := left + (n / 2)
	maxLenLeft := maxLenPosProductSubarray(nums, left, m)
	maxLenRight := maxLenPosProductSubarray(nums, m, right)
	productCrossLeft := 1
	posProdCrossLeft, negProdCrossLeft := 0, 0
	posProdCrossLeftIndex, negProdCrossLeftIndex := -1, -1
	for i := m - 1; i >= left; i-- {
		productCrossLeft *= nums[i]
		if productCrossLeft == 0 {
			break
		} else if productCrossLeft > 0 {
			posProdCrossLeft = productCrossLeft
			posProdCrossLeftIndex = i
		}
		if productCrossLeft < 0 {
			negProdCrossLeft = productCrossLeft
			negProdCrossLeftIndex = i
		}
	}
	productCrossRight := 1
	posProdCrossRight, negProdCrossRight := 0, 0
	posProdCrossRightIndex, negProdCrossRightIndex := -1, -1
	for i := m; i < right; i++ {
		productCrossRight *= nums[i]
		if productCrossRight == 0 {
			break
		}
		if productCrossRight > 0 {
			posProdCrossRight = productCrossRight
			posProdCrossRightIndex = i
		}
		if productCrossRight < 0 {
			negProdCrossRight = productCrossRight
			negProdCrossRightIndex = i
		}
	}
	maxLenPosPosCross := SubarrayMarkers{
		val:   posProdCrossLeft * posProdCrossRight,
		left:  posProdCrossLeftIndex,
		right: posProdCrossRightIndex + 1,
	}
	maxLenNegNegCross := SubarrayMarkers{
		val:   negProdCrossLeft * negProdCrossRight,
		left:  negProdCrossLeftIndex,
		right: negProdCrossRightIndex + 1,
	}
	return MaxLenPosSubarrayMarkers(maxLenLeft, maxLenRight, maxLenPosPosCross, maxLenNegNegCross)
}
