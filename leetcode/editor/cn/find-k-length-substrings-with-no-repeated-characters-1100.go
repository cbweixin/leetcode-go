package main

func numKLenSubstrNoRepeats(s string, k int) int {
	l, lookup := len(s), make(map[byte]int)
	if k > l {
		return 0
	}

	res := 0
	for i := 0; i < l; i++ {
		if _, ok := lookup[s[i]]; !ok {
			lookup[s[i]] = 1
		} else {
			lookup[s[i]] += 1
		}
		if i >= k {
			lookup[s[i-k]] -= 1
			if lookup[s[i-k]] == 0 {
				delete(lookup, s[i-k])
			}
		}
		if len(lookup) == k {
			res++
		}
	}

	return res

}
