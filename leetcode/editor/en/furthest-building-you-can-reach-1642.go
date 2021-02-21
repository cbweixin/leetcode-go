package main

import (
	"fmt"
	"sort"
)

//You are given an integer array heights representing the heights of buildings,
//some bricks, and some ladders.
//
// You start your journey from building 0 and move to the next building by possi
//bly using bricks or ladders.
//
// While moving from building i to building i+1 (0-indexed),
//
//
// If the current building's height is greater than or equal to the next buildin
//g's height, you do not need a ladder or bricks.
// If the current building's height is less than the next building's height, you
// can either use one ladder or (h[i+1] - h[i]) bricks.
//
//
// Return the furthest building index (0-indexed) you can reach if you use the g
//iven ladders and bricks optimally.
//
//
// Example 1:
//
//
//Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
//Output: 4
//Explanation: Starting at building 0, you can follow these steps:
//- Go to building 1 without using ladders nor bricks since 4 >= 2.
//- Go to building 2 using 5 bricks. You must use either bricks or ladders becau
//se 2 < 7.
//- Go to building 3 without using ladders nor bricks since 7 >= 6.
//- Go to building 4 using your only ladder. You must use either bricks or ladde
//rs because 6 < 9.
//It is impossible to go beyond building 4 because you do not have any more bric
//ks or ladders.
//
//
// Example 2:
//
//
//Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
//Output: 7
//
//
// Example 3:
//
//
//Input: heights = [14,3,19,3], bricks = 17, ladders = 0
//Output: 3
//
//
//
// Constraints:
//
//
// 1 <= heights.length <= 105
// 1 <= heights[i] <= 106
// 0 <= bricks <= 109
// 0 <= ladders <= heights.length
//
// Related Topics Binary Search Heap
// 👍 424 👎 29

// 2021-02-21 12:39:50

//leetcode submit region begin(Prohibit modification and deletion)
func furthestBuilding(heights []int, bricks int, ladders int) int {
	l := len(heights)
	left, right := 0, l-1

	can_move := func(idx int) bool {
		temp := []int{}

		for i := 0; i < idx; i++ {
			d := heights[i+1] - heights[i]
			if d > 0 {
				temp = append(temp, d)
			}
		}
		sort.Ints(temp)
		s := 0

		for i := 0; i < len(temp)-ladders; i++ {
			s += temp[i]
			if s > bricks {
				return false
			}
		}

		return true

	}

	for left <= right {
		mid := (left + right) >> 1
		if can_move(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {

	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks := 10
	ladders := 2

	fmt.Println(furthestBuilding(heights, bricks, ladders))

	heights = []int{4, 2, 7, 6, 9, 14, 12}
	bricks = 5
	ladders = 1
	fmt.Println(furthestBuilding(heights, bricks, ladders))

	heights = []int{17, 16, 5, 10, 10, 14, 7}
	bricks = 74
	ladders = 6
	fmt.Println(furthestBuilding(heights, bricks, ladders))
}
