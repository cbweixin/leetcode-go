package main

import (
	"fmt"
	"sort"
)

type Pair1 struct {
	I, J int
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	// knowledge, how to make slice as value for map
	buckets := make(map[int][]Pair1)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}
	dist := func(worker, bike []int) int {
		return abs(worker[0]-bike[0]) + abs(worker[1]-bike[1])
	}

	for i, w := range workers {
		for j, b := range bikes {
			buckets[dist(w, b)] = append(buckets[dist(w, b)], Pair1{i, j})
		}
	}

	used, res := make(map[int]bool), make([]int, len(workers))

	for i := 0; i < len(workers); i++ {
		res[i] = -1
	}
	// knowledge, how to sort map keys, compare with python, so verbose
	keys := make([]int, 0, len(buckets))
	for key := range buckets {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, d := range keys {
		for _, t := range buckets[d] {
			if res[t.I] == -1 && !used[t.J] {
				res[t.I] = t.J
				used[t.J] = true
			}
		}

	}

	return res

}

func main() {
	workers, bikes := [][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}}
	fmt.Println(assignBikes(workers, bikes))

}
