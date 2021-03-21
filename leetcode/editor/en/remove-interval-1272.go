package main

import "fmt"

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	a, b := toBeRemoved[0], toBeRemoved[1]
	//res := [][]int{}
	res := make([][]int, 0)

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for _, v := range intervals {
		s, e := v[0], v[1]
		if s < min(e, a) {
			res = append(res, []int{s, min(e, a)})
		}
		if e > max(b, s) {
			res = append(res, []int{max(b, s), e})
		}

	}

	return res

}
func main() {
	fmt.Println(removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}))
}
