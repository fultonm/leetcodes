package uglynumberii

func nthUglyNumber(n int) int {
	seive := generatePrimeSeive(1680)
	uglyMultiples := map[int]bool{2: true, 3: true, 5: true}
	uglyNumbers := make([]int, 1)
	uglyNumbers[0] = 1
OUTER:
	for len(uglyNumbers) < n {
		i := uglyNumbers[len(uglyNumbers)-1] + 1
		for {
			if isUglyNumber(seive[i], uglyMultiples) {
				uglyNumbers = append(uglyNumbers, i)
				continue OUTER
			}
		}

	}
	return uglyNumbers[n-1]
}

func isUglyNumber(factors []int, uglyMultiples map[int]bool) bool {
	if len(factors) == 0 {
		return true
	}
	if len(factors) > 3 {
		return false
	}
	for _, f := range factors {
		if _, ok := uglyMultiples[f]; !ok {
			return false
		}
	}
	return true
}

func generatePrimeSeive(n int) map[int][]int {
	seive := make(map[int][]int)
	for i := 2; i <= n; i++ {
		multiple := i
		for multiple <= n {
			seive[multiple] = append(seive[multiple], i)
			multiple += i
		}
	}
	return seive
}
