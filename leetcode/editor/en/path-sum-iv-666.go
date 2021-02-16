package main

import "fmt"

type Key3 struct {
	Level int
	Index int
}

func pathSum(nums []int) int {
	l := len(nums)
	lookup := make(map[Key3]int)

	res := 0
	for i := l - 1; i >= 0; i-- {
		level, idx, val := int(nums[i]/100), int(nums[i]%100/10), nums[i]%10
		key, parent := Key3{level, idx}, Key3{level - 1, (idx + 1) >> 1}
		if _, ok := lookup[key]; !ok {
			lookup[key] = 1
		}
		res += val * lookup[key]
		if _, ok := lookup[parent]; !ok {
			lookup[parent] = 0
		}
		lookup[parent] += lookup[key]

	}

	return res

}

func main() {

	nums := []int{113, 215, 221}
	fmt.Println(pathSum(nums))

}
