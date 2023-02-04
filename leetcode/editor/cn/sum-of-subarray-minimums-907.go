package main

import "fmt"

//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//
// 由于答案可能很大，因此 返回答案模 10^9 + 7 。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
//
// 示例 2：
//
//
//输入：arr = [11,81,94,43,3]
//输出：444
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 3 * 10⁴
// 1 <= arr[i] <= 3 * 10⁴
//
//
//
//
// Related Topics 栈 数组 动态规划 单调栈 👍 614 👎 0

//2023-02-04 11:11:32

// leetcode submit region begin(Prohibit modification and deletion)
func sumSubarrayMins(arr []int) int {
	st := make([]int, 0)
	tArr := []int{0}
	tArr = append(tArr, arr...)
	tArr = append(tArr, 0)
	mod, res := 1000000007, 0
	for i, v := range tArr {
		for len(st) > 0 && tArr[st[len(st)-1]] > v {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			k := st[len(st)-1]
			res += (i - j) * (j - k) * tArr[j] % mod
			res %= mod
		}
		st = append(st, i)
	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}
