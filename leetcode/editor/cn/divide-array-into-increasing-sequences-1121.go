package main

func canDivideIntoSubsequences(nums []int, k int) bool {
	counter := make(map[int]int)

	maxOccurance := 0
	for _, v := range nums {
		if _, ok := counter[v]; !ok {
			counter[v] = 1
		} else {
			counter[v] += 1
			if maxOccurance < counter[v] {
				maxOccurance = counter[v]
			}
		}
	}
	return len(nums) >= maxOccurance*k

}
