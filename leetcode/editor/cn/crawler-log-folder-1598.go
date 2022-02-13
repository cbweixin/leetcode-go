package main

// 每当用户执行变更文件夹操作时，LeetCode 文件系统都会保存一条日志记录。
//
// 下面给出对变更操作的说明：
//
//
// "../" ：移动到当前文件夹的父文件夹。如果已经在主文件夹下，则 继续停留在当前文件夹 。
// "./" ：继续停留在当前文件夹。
// "x/" ：移动到名为 x 的子文件夹中。题目数据 保证总是存在文件夹 x 。
//
//
// 给你一个字符串列表 logs ，其中 logs[i] 是用户在 iᵗʰ 步执行的操作。
//
// 文件系统启动时位于主文件夹，然后执行 logs 中的操作。
//
// 执行完所有变更文件夹操作后，请你找出 返回主文件夹所需的最小步数 。
//
//
//
// 示例 1：
//
//
//
// 输入：logs = ["d1/","d2/","../","d21/","./"]
// 输出：2
// 解释：执行 "../" 操作变更文件夹 2 次，即可回到主文件夹
//
//
// 示例 2：
//
//
//
// 输入：logs = ["d1/","d2/","./","d3/","../","d31/"]
// 输出：3
//
//
// 示例 3：
//
// 输入：logs = ["d1/","../","../","../"]
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= logs.length <= 10³
// 2 <= logs[i].length <= 10
// logs[i] 包含小写英文字母，数字，'.' 和 '/'
// logs[i] 符合语句中描述的格式
// 文件夹名称由小写英文字母和数字组成
//
// Related Topics 栈 数组 字符串 👍 18 👎 0

// 2022-02-13 13:32:45
// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(logs []string) int {
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, v := range logs {
		if v == "../" {
			res = max(0, res-1)
		} else if v != "./" {
			res += 1
		}
	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
