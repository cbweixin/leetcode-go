package main

import (
	"fmt"
	"math"
	"strings"
)

//There is a box protected by a password. The password is a sequence of n digits
// where each digit can be one of the first k digits 0, 1, ..., k-1.
//
// While entering a password, the last n digits entered will automatically be ma
//tched against the correct password.
//
// For example, assuming the correct password is "345", if you type "012345", th
//e box will open because the correct password matches the suffix of the entered p
//assword.
//
// Return any password of minimum length that is guaranteed to open the box at s
//ome point of entering it.
//
//
//
// Example 1:
//
//
//Input: n = 1, k = 2
//Output: "01"
//Note: "10" will be accepted too.
//
//
// Example 2:
//
//
//Input: n = 2, k = 2
//Output: "00110"
//Note: "01100", "10011", "11001" will be accepted too.
//
//
//
//
// Note:
//
//
// n will be in the range [1, 4].
// k will be in the range [1, 10].
// k^n will be at most 4096.
//
//
//
// Related Topics Math Depth-first Search
// 👍 516 👎 753

// 2021-01-22 12:45:39

//leetcode submit region begin(Prohibit modification and deletion)
func crackSafe(n int, k int) string {
	res, set := strings.Repeat("0", n), make(map[string]bool)
	set[res] = true

	for i := 0; i < int(math.Pow(float64(k), float64(n))); i++ {
		l := len(res)
		pre := res[l-n+1 : l]
		for j := k - 1; j >= 0; j-- {
			t := pre + string(j+'0')
			if _, ok := set[t]; !ok {
				set[t] = true
				res += string(j + '0')
				break
			}
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(crackSafe(2, 2))
}
