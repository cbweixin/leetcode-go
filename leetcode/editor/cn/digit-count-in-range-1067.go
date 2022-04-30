package main

func digitsCount(d int, low int, high int) int {

	countDigit := func(n int) int {
		k, res := 1, 0
		for k = 1; k <= n; k *= 10 {
			r, m := n%k, n/k
			e, xyz := m%10, m/10
			res += xyz * k
			if d == 0 {
				res -= k
			}
			if e > d {
				res += k
			} else if e == d {
				res += 1 + r
			}
		}
		return res
	}
	return countDigit(high) - countDigit(low-1)
}
