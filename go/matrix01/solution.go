package matrix01

func UpdateMatrix(mat [][]int) [][]int {
	return updateMatrix(mat)
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	q := []Point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				mat[i][j] = IntMax
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

	for len(q) > 0 {
		cur := poll(&q)
		for _, e := range dir {
			neighbor := Point{
				cur.y + e.y,
				cur.x + e.x,
			}
			if neighbor.y < 0 || neighbor.y >= m || neighbor.x < 0 || neighbor.x >= n || mat[neighbor.y][neighbor.x] <= mat[cur.y][cur.x] {
				continue
			}
			mat[neighbor.y][neighbor.x] = mat[cur.y][cur.x] + 1
			offer(&q, neighbor)
		}

	}

	return mat
}

type Point struct {
	y int
	x int
}

func poll(q *[]Point) Point {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func offer(q *[]Point, p Point) {
	*q = append(*q, p)
}

func MinInt(args ...int) int {
	r := IntMax
	for _, e := range args {
		if e < r {
			r = e
		}
	}
	return r
}

const IntMax = int(^uint(0) >> 1)
