package main

func orangesRotting(grid [][]int) int {
	/** vv Types and functions vv */
	type Point struct {
		y int
		x int
	}
	poll := func(q *[]Point) Point {
		p := (*q)[0]
		*q = (*q)[1:]
		return p
	}
	offer := func(q *[]Point, p Point) {
		*q = append(*q, p)
	}

	const IntMax = int(^uint(0) >> 1)
	const IntMin = -int(^uint(0)>>1) - 1

	MaxInt := func(args ...int) int {
		r := IntMin
		for _, e := range args {
			if e > r {
				r = e
			}
		}
		return r
	}
	/** ^^ Types and functions ^^ */
	m, n := len(grid), len(grid[0])
	q := []Point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = IntMax
			} else if grid[i][j] == 2 {
				grid[i][j] = 0
				offer(&q, Point{i, j})
			} else {
				grid[i][j] = IntMin
			}

		}
	}
	dir := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	max := 0
	for len(q) > 0 {
		cur := poll(&q)
		for _, e := range dir {
			neighbor := Point{
				cur.y + e.y,
				cur.x + e.x,
			}
			if neighbor.y < 0 || neighbor.y >= m || neighbor.x < 0 || neighbor.x >= n ||
				grid[neighbor.y][neighbor.x] < 0 || grid[neighbor.y][neighbor.x] <= grid[cur.y][cur.x] {
				continue
			}
			grid[neighbor.y][neighbor.x] = grid[cur.y][cur.x] + 1
			max = MaxInt(max, grid[neighbor.y][neighbor.x])
			offer(&q, neighbor)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == IntMax {
				return -1
			}
		}
	}
	return max
}
