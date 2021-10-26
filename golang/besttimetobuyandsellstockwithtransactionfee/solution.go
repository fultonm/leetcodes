package besttimetobuyandsellstockwithtransactionfee

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, 2)
	}
	memo[0][0] = 0
	memo[0][1] = -prices[0]
	memo[1][0] = MaxInt(0, prices[1]-prices[0]-fee)
	memo[1][1] = MaxInt(-prices[0], -prices[1])
	for i := 2; i < n; i++ {
		memo[i][0] = MaxInt(memo[i-1][0], memo[i-1][1]+prices[i]-fee)
		memo[i][1] = MaxInt(memo[i-1][0]-prices[i], memo[i-1][1])
	}
	return memo[n-1][0]
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
