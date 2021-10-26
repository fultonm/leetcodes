package wordbreak

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	memo := make([]bool, n+1)
	memo[n] = true

	for i := n - 1; i >= 0; i-- {
	OUTER:
		for j := i + 1; j <= n; j++ {
			sub := s[i:j]
			for _, e := range wordDict {
				if sub == e {
					memo[i] = memo[i+len(e)]
					if memo[i] {
						break OUTER
					}
				}
			}
		}
	}
	return memo[0]
}
