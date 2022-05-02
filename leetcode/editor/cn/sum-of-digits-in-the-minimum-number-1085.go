package main

func sumOfDigits(nums []int) int {
	min := 110
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return 1 - (min%10+min/10)&1

}
