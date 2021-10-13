package main

func letterCasePermutation(s string) []string {
	bs := []byte(s)
	n := len(bs)
	r := []string{}
	loopLetterCase(0, n, &bs, &r)
	return r
}

func loopLetterCase(l int, n int, bs *[]byte, r *[]string) {
	if l < n {
		loopLetterCase(l+1, n, bs, r)
		if (*bs)[l] >= 'a' && (*bs)[l] <= 'z' {
			tmp := (*bs)[l]
			(*bs)[l] = tmp - 32
			loopLetterCase(l+1, n, bs, r)
			(*bs)[l] = tmp
		} else if (*bs)[l] >= 'A' && (*bs)[l] <= 'Z' {
			tmp := (*bs)[l]
			(*bs)[l] = tmp + 32
			loopLetterCase(l+1, n, bs, r)
			(*bs)[l] = tmp
		}
	} else {
		newPerm := make([]byte, n)
		copy(newPerm, *bs)
		*r = append(*r, string(newPerm))
	}
}
