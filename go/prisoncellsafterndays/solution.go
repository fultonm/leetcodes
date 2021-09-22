package prisoncellsafterndays

import (
	"fmt"
)

func prisonAfterNDays(cells []int, n int) []int {
	if n <= 0 {
		return cells
	}
	cellBits := cellsToBitstring(cells)
	memo := make(map[int][]int)
	memo[cellBits] = append(memo[cellBits], n)
	for j := 0; j < n; j++ {
		newCellBits := 2
		for i := 0; i < len(cells)-1; i++ {
			if (cellBits>>len(cells)-2-i)&1 == (cellBits>>len(cells)-i)&1 {
				newCellBits += 1
			}
			newCellBits <<= 1
		}
		fmt.Println(newCellBits)
		cellBits = newCellBits
		memo[cellBits] = append(memo[cellBits], j)
		if len(memo[cellBits]) > 1 {
			a := memo[cellBits][0]
			b := memo[cellBits][1]
			// return the right cells
			fmt.Printf("Pattern deteected %d %d\n", a, b)
		}
	}
	return bitstringToCells(cellBits)
}

func cellsToBitstring(cells []int) int {
	r := 0
	for i := 0; i < len(cells); i++ {
		if cells[i] == 1 {
			r += intPow(2, len(cells)-i-1)
		}
	}
	return r
}

func bitstringToCells(b int) []int {
	c := make([]int, 9)
	i := 8
	for b > 0 {
		if b%2 == 1 {
			c[i] = 1
		}
		i--
		b = b >> 1
	}
	return c
}

func intPow(base int, exp int) int {
	if exp < -1 {
		return 0
	}
	if exp == 0 || exp == -1 {
		return 1
	}
	r := base
	for i := 1; i < exp; i++ {
		r *= base
	}
	return r
}
