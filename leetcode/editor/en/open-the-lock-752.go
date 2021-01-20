package main

import (
	"fmt"
	"time"
)

//You have a lock in front of you with 4 circular wheels. Each wheel has 10 slot
//s: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freel
//y and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each
// move consists of turning one wheel one slot.
//
// The lock initially starts at '0000', a string representing the state of the 4
// wheels.
//
// You are given a list of deadends dead ends, meaning if the lock displays any
//of these codes, the wheels of the lock will stop turning and you will be unable
//to open it.
//
// Given a target representing the value of the wheels that will unlock the lock
//, return the minimum total number of turns required to open the lock, or -1 if i
//t is impossible.
//
//
// Example 1:
//
//
//Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//Output: 6
//Explanation:
//A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "12
//01" -> "1202" -> "0202".
//Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would
// be invalid,
//because the wheels of the lock become stuck after the display becomes the dead
// end "0102".
//
//
// Example 2:
//
//
//Input: deadends = ["8888"], target = "0009"
//Output: 1
//Explanation:
//We can turn the last wheel in reverse to move from "0000" -> "0009".
//
//
// Example 3:
//
//
//Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], t
//arget = "8888"
//Output: -1
//Explanation:
//We can't reach the target without getting stuck.
//
//
// Example 4:
//
//
//Input: deadends = ["0000"], target = "8888"
//Output: -1
//
//
//
// Constraints:
//
//
// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target will not be in the list deadends.
// target and deadends[i] consist of digits only.
//
// Related Topics Breadth-first Search
// 👍 1390 👎 54

// 2021-01-20 10:40:43

//leetcode submit region begin(Prohibit modification and deletion)
func openLock(deadends []string, target string) int {
	lookup := make(map[string]bool)
	for _, v := range deadends {
		lookup[v] = true
	}
	if lookup[target] || lookup["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	seen := make(map[string]bool)
	que := []string{}
	que = append(que, "0000")

	// review, how to do bfs in golang
	for time := 1; len(que) != 0; time++ {
		for n := len(que); n > 0; n-- {
			cur := que[0]
			que = que[1:]
			for i := 0; i < 4; i++ {
				for j := -1; j < 2; j++ {
					// review, rune is character in golang
					s := []rune(cur)
					// review, there is a big differences between go and python, -1 % 10 == -1 for go, but 9 for python
					// that's the reason why we need to + 10 again
					s[i] = rune((int(s[i])-int('0')+j+10)%10 + int('0'))
					if string(s) == target {
						return time
					}
					if !lookup[string(s)] && !seen[string(s)] {
						que = append(que, string(s))
						seen[string(s)] = true
					}

				}
			}

		}

	}

	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	start := time.Now()
	//res := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	res := openLock([]string{"8888"}, "0009")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Milliseconds(), "ms")
	fmt.Println(res)

}
