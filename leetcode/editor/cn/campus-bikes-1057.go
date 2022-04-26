package main

import (
	"sort"
)

type Pair struct {
	I, J int
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dist := func(w, b []int) int {
		return abs(w[0]-b[0]) + abs(w[1]-b[1])
	}
	dists := make(map[int][]Pair)

	for i, w := range workers {
		for j, b := range bikes {
			dists[dist(w, b)] = append(dists[dist(w, b)], Pair{i, j})
		}
	}
	m, n := len(workers), len(bikes)
	res, used := make([]int, m), make(map[int]bool)
	for i := 0; i < m; i++ {
		res[i] = -1
	}

	keys := make([]int, n)
	for key := range dists {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, d := range keys {
		for _, t := range dists[d] {
			if res[t.I] == -1 && used[t.J] == false {
				res[t.I] = t.J
				used[t.J] = true
			}
		}
	}

	return res

}
