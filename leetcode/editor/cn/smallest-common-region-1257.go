package main

//给你一些区域列表 regions ，每个列表的第一个区域都包含这个列表内所有其他区域。
//
// 很自然地，如果区域 X 包含区域 Y ，那么区域 X 比区域 Y 大。
//
// 给定两个区域 region1 和 region2 ，找到同时包含这两个区域的 最小 区域。
//
// 如果区域列表中 r1 包含 r2 和 r3 ，那么数据保证 r2 不会包含 r3 。
//
// 数据同样保证最小公共区域一定存在。
//
//
//
// 示例 1：
//
//
//输入：
//regions = [["Earth","North America","South America"],
//["North America","United States","Canada"],
//["United States","New York","Boston"],
//["Canada","Ontario","Quebec"],
//["South America","Brazil"]],
//region1 = "Quebec",
//region2 = "New York"
//输出："North America"
//
//
//
//
// 提示：
//
//
// 2 <= regions.length <= 10^4
// region1 != region2
// 所有字符串只包含英文字母和空格，且最多只有 20 个字母。
//
// Related Topics 树 深度优先搜索 广度优先搜索 数组 哈希表 字符串 👍 42 👎 0

// 2022-06-18 12:50:15
//leetcode submit region begin(Prohibit modification and deletion)
func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	parents, pset := make(map[string]string), make(map[string]bool)

	for _, rs := range regions {
		for i := 1; i < len(rs); i++ {
			parents[rs[i]] = rs[0]
		}
	}

	ok := true
	for ok {
		pset[region1] = true
		region1, ok = parents[region1]
	}
	ok = true
	for ok {
		if _, ook := pset[region2]; ook {
			return region2
		}
		region2, ok = parents[region2]
	}
	return ""

}

//leetcode submit region end(Prohibit modification and deletion)
