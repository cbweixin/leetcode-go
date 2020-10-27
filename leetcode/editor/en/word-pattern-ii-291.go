package main

func wordPatternMatch(pattern string, str string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)

	// review , how to define a recursive function
	var dfs func(int, int) bool
	dfs = func(p int, s int) bool {
		if p == len(pattern) && s == len(str) {
			return true
		}
		if p == len(pattern) || s == len(str) {
			return false
		}
		for i := s; i < len(str); i++ {
			w, c := str[s:i+1], pattern[p]
			_, ok1 := s2p[w]
			v2, ok2 := p2s[c]

			if !ok1 && !ok2 {
				p2s[c] = w
				s2p[w] = c
				// call recursivly
				if dfs(p+1, i+1) {
					return true
				}
				// backtrack
				delete(p2s, c)
				delete(s2p, w)
			} else if ok2 && v2 == w {
				return dfs(p+1, i+1)
			}
		}

		return false

	}

	return dfs(0, 0)

}

func main() {

}
