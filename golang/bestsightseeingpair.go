package main

func maxScoreSightseeingPair(values []int) int {
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
	// MinInt := func(args ...int) int {
	// 	if len(args) == 0 {
	// 		return IntMin
	// 	}
	// 	r := IntMax
	// 	for _, e := range args {
	// 		if e < r {
	// 			r = e
	// 		}
	// 	}
	// 	return r
	// }
	getScore := func(values []int, left int, right int) int {
		if left >= right {
			return -1
		}
		return values[left] + values[right] + left - right
	}
	left, right, next := 0, 1, 2
	score := getScore(values, left, right)
	n := len(values)
	for i := 2; i < n; i++ {
		scoreLeft := getScore(values, left, i)
		scoreRight := getScore(values, right, i)
		scoreNext := getScore(values, next, i)
		scoreMax := MaxInt(scoreLeft, scoreRight, scoreNext)
		if scoreMax > score {
			switch scoreMax {
			case scoreLeft:
				right = i
				score = scoreLeft
			case scoreRight:
				left = right
				right = i
				score = scoreRight
			case scoreNext:
				left = next
				right = i
				score = scoreNext
			}
			next = i + 1
		} else if values[i] > values[next]+next-i {
			next = i
		}
	}
	return score
}
