package main

import "fmt"

//给你两个数组 arr1 和 arr2 ，它们一开始都是空的。你需要往它们中添加正整数，使它们满足以下条件：
//
//
// arr1 包含 uniqueCnt1 个 互不相同 的正整数，每个整数都 不能 被 divisor1 整除 。
// arr2 包含 uniqueCnt2 个 互不相同 的正整数，每个整数都 不能 被 divisor2 整除 。
// arr1 和 arr2 中的元素 互不相同 。
//
//
// 给你 divisor1 ，divisor2 ，uniqueCnt1 和 uniqueCnt2 ，请你返回两个数组中 最大元素 的 最小值 。
//
//
//
// 示例 1：
//
//
//输入：divisor1 = 2, divisor2 = 7, uniqueCnt1 = 1, uniqueCnt2 = 3
//输出：4
//解释：
//我们可以把前 4 个自然数划分到 arr1 和 arr2 中。
//arr1 = [1] 和 arr2 = [2,3,4] 。
//可以看出两个数组都满足条件。
//最大值是 4 ，所以返回 4 。
//
//
// 示例 2：
//
//
//输入：divisor1 = 3, divisor2 = 5, uniqueCnt1 = 2, uniqueCnt2 = 1
//输出：3
//解释：
//arr1 = [1,2] 和 arr2 = [3] 满足所有条件。
//最大值是 3 ，所以返回 3 。
//
// 示例 3：
//
//
//输入：divisor1 = 2, divisor2 = 4, uniqueCnt1 = 8, uniqueCnt2 = 2
//输出：15
//解释：
//最终数组为 arr1 = [1,3,5,7,9,11,13,15] 和 arr2 = [2,6] 。
//上述方案是满足所有条件的最优解。
//
//
//
//
// 提示：
//
//
// 2 <= divisor1, divisor2 <= 10⁵
// 1 <= uniqueCnt1, uniqueCnt2 < 10⁹
// 2 <= uniqueCnt1 + uniqueCnt2 <= 10⁹
//
//
// Related Topics 数学 二分查找 数论 👍 29 👎 0

//2023-05-09 19:23:28

// leetcode submit region begin(Prohibit modification and deletion)
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {

	i := 1
	a, b := uniqueCnt1, uniqueCnt2

	for true {
		if uniqueCnt1 > 0 && i%divisor1 != 0 {
			uniqueCnt1--
		} else if uniqueCnt2 > 0 && i%divisor2 != 0 {
			uniqueCnt2--
		}

		if uniqueCnt2 == 0 && uniqueCnt1 == 0 {
			break
		}
		i++
	}

	j := 1
	uniqueCnt1, uniqueCnt2 = a, b
	for true {
		if uniqueCnt2 > 0 && j%divisor2 != 0 {
			uniqueCnt2--
		} else if uniqueCnt1 > 0 && j%divisor1 != 0 {
			uniqueCnt1--
		}

		if uniqueCnt1 == 0 && uniqueCnt2 == 0 {
			break
		}
		j++
	}

	if i > j {
		return j
	}

	return i

}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minimizeSet(321, 734, 1234, 2341))
}
