package main

func orangesRotting(grid [][]int) int {
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

	m, n := len(grid), len(grid[0])
	q := []Point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				grid[i][j] = IntMax
			} else {
				offer(&q, Point{i, j})
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
				grid[neighbor.y][neighbor.x] == 0 || grid[neighbor.y][neighbor.x] <= grid[cur.y][cur.x] {
				continue
			}
			grid[neighbor.y][neighbor.x] = grid[cur.y][cur.x] + 1
			offer(&q, neighbor)
		}

	}

	return max
}
