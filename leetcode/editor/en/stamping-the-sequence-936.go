package main

import (
	"fmt"
	"strings"
)

//You want to form a target string of lowercase letters.
//
// At the beginning, your sequence is target.length '?' marks. You also have a s
//tamp of lowercase letters.
//
// On each turn, you may place the stamp over the sequence, and replace every le
//tter in the sequence with the corresponding letter from the stamp. You can make
//up to 10 * target.length turns.
//
// For example, if the initial sequence is "?????", and your stamp is "abc", the
//n you may make "abc??", "?abc?", "??abc" in the first turn. (Note that the stamp
// must be fully contained in the boundaries of the sequence in order to stamp.)
//
// If the sequence is possible to stamp, then return an array of the index of th
//e left-most letter being stamped at each turn. If the sequence is not possible t
//o stamp, return an empty array.
//
// For example, if the sequence is "ababc", and the stamp is "abc", then we coul
//d return the answer [0, 2], corresponding to the moves "?????" -> "abc??" -> "ab
//abc".
//
// Also, if the sequence is possible to stamp, it is guaranteed it is possible t
//o stamp within 10 * target.length moves. Any answers specifying more than this n
//umber of moves will not be accepted.
//
//
//
// Example 1:
//
//
//Input: stamp = "abc", target = "ababc"
//Output: [0,2]
//([1,0,2] would also be accepted as an answer, as well as some other answers.)
//
//
//
// Example 2:
//
//
//Input: stamp = "abca", target = "aabcaca"
//Output: [3,0,1]
//
//
//
//
//
// Note:
//
//
//
//
// 1 <= stamp.length <= target.length <= 1000
// stamp and target only contain lowercase letters.
// Related Topics String Greedy
// 👍 240 👎 59

// 2021-02-03 09:54:52

//leetcode submit region begin(Prohibit modification and deletion)
func movesToStamp(stamp string, target string) []int {
	L, M, res := len(stamp), len(target), []int{}
	can_change := func(i int) bool {
		changed := false
		for j := 0; j < L; j++ {
			if target[i+j] == '?' {
				continue
			}
			if target[i+j] != stamp[j] {
				return false
			}
			changed = true
		}
		if changed {
			target = strings.Replace(target, target[i:i+L], strings.Repeat("?", L), 1)
			res = append(res, i)
		}

		return changed
	}

	reverse := func(arr []int) {
		i, j := 0, len(arr)-1

		for i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	need_change := true

	for need_change {
		need_change = false
		for i := 0; i < M-L+1; i++ {
			need_change = need_change || can_change(i)
		}
	}
	if target == strings.Repeat("?", M) {
		reverse(res)
		return res
	} else {
		return []int{}

	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(movesToStamp("abc", "ababc"))
}
