package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
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
	getProfit := func(left int, right int) int {
		if left >= right {
			return -1
		}
		return prices[right] - prices[left]
	}
	left, right, next := 0, 1, 1
	profit := getProfit(left, right)
	n := len(prices)
	for i := 2; i < n; i++ {
		profitLeft := getProfit(left, i)
		profitRight := getProfit(right, i)
		profitNext := getProfit(next, i)
		profitMax := MaxInt(profitLeft, profitRight, profitNext)
		if profitMax > profit {
			switch profitMax {
			case profitLeft:
				right = i
				profit = profitLeft
			case profitRight:
				left = right
				right = i
				profit = profitRight
			case profitNext:
				left = next
				right = i
				profit = profitNext
			}
			next = i + 1
		} else if prices[i] < prices[next] {
			next = i
		}
	}
	return MaxInt(0, profit)

}
