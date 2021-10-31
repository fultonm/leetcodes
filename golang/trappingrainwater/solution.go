package trappingrainwater

type Cavity struct {
	left  int
	right int
}

func trap(height []int) int {
	memo := make([]Cavity, 0)
	n := len(height)
	left := 0
	for i := 0; i < n; i++ {
		if height[left] > height[i] {
			continue
		}
		if height[left] < height[i] {
			for j := len(memo) - 1; j >= 0; j-- {
				if height[memo[j].left] >= height[i] {
					left = memo[j].left
					memo = memo[:j]
				}
			}
		}
		memo = append(memo, Cavity{left, i})
		left = i
	}
	return 0
}

func isLocalMaxima(i int, height []int) bool {
	if i == 0 && len(height) > 1 {
		return height[i] > height[i+1]
	}
	if i == len(height)-1 && len(height) > 1 {
		return height[i] > height[i-1]
	}
	if height[i] > height[i-1] && height[i] >= height[i+1] {
		return true
	}
	return false
	// [1234], any false
	// [3212], 2 true
	// [421223], 3 true
	// [12322], 4 true ??
	// [3433241234543443232125]

}
