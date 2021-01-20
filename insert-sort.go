package main

import (
	"fmt"
	"sort"
)

func InsertSorted(s []int, e int) []int {
	s = append(s, 0)
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 6, 7, 8, 9}
	s = InsertSorted(s, 5)
	fmt.Println(s)
}
