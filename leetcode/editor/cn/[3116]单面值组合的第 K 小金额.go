package main

import "sort"

//给你一个整数数组 coins 表示不同面额的硬币，另给你一个整数 k 。
//
// 你有无限量的每种面额的硬币。但是，你 不能 组合使用不同面额的硬币。
//
// 返回使用这些硬币能制造的 第 kᵗʰ 小 金额。
//
//
//
// 示例 1：
//
//
// 输入： coins = [3,6,9], k = 3
//
//
// 输出： 9
//
// 解释：给定的硬币可以制造以下金额： 3元硬币产生3的倍数：3, 6, 9, 12, 15等。 6元硬币产生6的倍数：6, 12, 18, 24等。 9元硬
//币产生9的倍数：9, 18, 27, 36等。 所有硬币合起来可以产生：3, 6, 9, 12, 15等。
//
// 示例 2：
//
//
// 输入：coins = [5,2], k = 7
//
//
// 输出：12
//
// 解释：给定的硬币可以制造以下金额： 5元硬币产生5的倍数：5, 10, 15, 20等。 2元硬币产生2的倍数：2, 4, 6, 8, 10, 12等。
//所有硬币合起来可以产生：2, 4, 5, 6, 8, 10, 12, 14, 15等。
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 15
// 1 <= coins[i] <= 25
// 1 <= k <= 2 * 10⁹
// coins 包含两两不同的整数。
//
//
// Related Topics 位运算 数组 数学 二分查找 组合数学 数论 👍 28 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findKthSmallest(coins []int, k int) int64 {
	sort.Ints(coins)
	base := make([]int64, 0, len(coins))

	for _, coin := range coins {
		c := int64(coin)
		isMultiple := false

		for _, b := range base {
			if c%b == 0 {
				isMultiple = true
				break
			}
		}
		if !isMultiple {
			base = append(base, c)
		}
	}

	n, lo, hi, pos, neg := len(base), int64(1), base[0]*int64(k), make([]int64, 0), make([]int64, 0)
	for mask := 1; mask < (1 << n); mask++ {
		lcmValue, bitLen, valid := int64(1), 0, true
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			coin := base[i]
			g := gcd(lcmValue, coin)

			if lcmValue > hi/(coin/g) {
				valid = false
				break
			}
			lcmValue = lcmValue / g * coin
			bitLen++
		}

		if !valid {
			continue
		}
		if bitLen&1 == 1 {
			pos = append(pos, lcmValue)
		} else {
			neg = append(neg, lcmValue)
		}

	}

	for lo < hi {
		mid := (lo + hi) / 2
		count := int64(0)

		for _, value := range pos {
			count += mid / value
		}
		for _, value := range neg {
			count -= mid / value
		}

		if count < int64(k) {
			lo = mid + 1

		} else {
			hi = mid
		}
	}
	return lo
}

//leetcode submit region end(Prohibit modification and deletion)
