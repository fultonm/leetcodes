package main

func minimumTotal(triangle [][]int) int {
	const IntMax = int(^uint(0) >> 1)
	minInt := func(args ...int) int {
		r := IntMax
		for _, e := range args {
			if e < r {
				r = e
			}
		}
		return r
	}
	h := len(triangle)
	for l := 1; l < h; l++ {
		w := len(triangle[l])
		for i := 0; i < w; i++ {
			if i == 0 {
				triangle[l][i] = triangle[l][i] + triangle[l-1][i]
			} else if i >= w-1 {
				triangle[l][i] = triangle[l][i] + triangle[l-1][i-1]
			} else {
				triangle[l][i] = minInt(triangle[l][i]+triangle[l-1][i],
					triangle[l][i]+triangle[l-1][i-1])
			}
		}
	}
	return minInt(triangle[h-1]...)
}
