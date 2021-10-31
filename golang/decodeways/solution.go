package decodeways

func numDecodings(s string) int {
	memo := make(map[int]int)
	i := len(s) - 1
	for i >= 0 {
		memo[i] = ways(i, s, memo)
		i--
	}
	return memo[0]
}

// 2101
func ways(i int, s string, memo map[int]int) int {
	n := len(s)
	if i >= n {
		return 1
	}
	if v, ok := memo[i]; ok {
		return v
	}
	if s[i] == '0' {
		return 0
	}
	if i+1 < n {
		if s[i] == '1' {
			if s[i+1] == '0' {
				return ways(i+2, s, memo)
			}
			if s[i+1] > '0' {
				return ways(i+1, s, memo) + +ways(i+2, s, memo)
			}
		}
		if s[i] == '2' {
			if s[i+1] == '0' {
				return ways(i+2, s, memo)
			}
			if s[i+1] > '0' && s[i+1] < '7' {
				return ways(i+1, s, memo) + ways(i+2, s, memo)
			}
		}
	}
	return ways(i+1, s, memo)
}
