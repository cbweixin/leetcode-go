package main

//给出一个二进制数组 data，你需要通过交换位置，将数组中 任何位置 上的 1 组合到一起，并返回所有可能中所需 最少的交换次数。
//
//
//
// 示例 1:
//
//
//输入: data = [1,0,1,0,1]
//输出: 1
//解释:
//有三种可能的方法可以把所有的 1 组合在一起：
//[1,1,1,0,0]，交换 1 次；
//[0,1,1,1,0]，交换 2 次；
//[0,0,1,1,1]，交换 1 次。
//所以最少的交换次数为 1。
//
//
// 示例 2:
//
//
//输入：data = [0,0,0,1,0]
//输出：0
//解释：
//由于数组中只有一个 1，所以不需要交换。
//
// 示例 3:
//
//
//输入：data = [1,0,1,0,1,0,0,1,1,0,1]
//输出：3
//解释：
//交换 3 次，一种可行的只用 3 次交换的解决方案是 [0,0,0,0,0,1,1,1,1,1,1]。
//
//
// 示例 4:
//
//
//输入: data = [1,0,1,0,1,0,1,1,1,0,1,0,0,1,1,1,0,0,1,1,1,0,1,0,1,1,0,0,0,1,1,1,1,
//0,0,1]
//输出: 8
//
//
//
//
// 提示:
//
//
// 1 <= data.length <= 10⁵
// data[i] == 0 or 1.
//
// Related Topics 数组 滑动窗口 👍 69 👎 0

// 2022-05-21 11:58:06
//leetcode submit region begin(Prohibit modification and deletion)
func minSwaps(data []int) int {
	sum := func(arr []int) int {
		s := 0
		for _, v := range arr {
			s += v
		}

		return s
	}
	max := func(a, b int) int {
		if a <= b {
			return b
		}
		return a
	}
	s, l := sum(data), len(data)
	t_s := sum(data[:s])
	l_s := t_s
	for i := 1; i+s-1 < l; i++ {
		l_s = l_s - data[i-1] + data[i+s-1]
		t_s = max(t_s, l_s)
	}
	return s - t_s
}

//leetcode submit region end(Prohibit modification and deletion)
