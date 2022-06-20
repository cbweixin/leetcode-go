package main

//给你一个近义词表 synonyms 和一个句子 text ， synonyms 表中是一些近义词对 ，你可以将句子 text 中每个单词用它的近义词来替换。
//
//
// 请你找出所有用近义词替换后的句子，按 字典序排序 后返回。
//
//
//
// 示例 1：
//
//
//输入：
//synonyms = [["happy","joy"],["sad","sorrow"],["joy","cheerful"]],
//text = "I am happy today but was sad yesterday"
//输出：
//["I am cheerful today but was sad yesterday",
//"I am cheerful today but was sorrow yesterday",
//"I am happy today but was sad yesterday",
//"I am happy today but was sorrow yesterday",
//"I am joy today but was sad yesterday",
//"I am joy today but was sorrow yesterday"]
//
//
//
//
// 提示：
//
//
// 0 <= synonyms.length <= 10
// synonyms[i].length == 2
// synonyms[0] != synonyms[1]
// 所有单词仅包含英文字母，且长度最多为 10 。
// text 最多包含 10 个单词，且单词间用单个空格分隔开。
//
// Related Topics 并查集 数组 哈希表 字符串 回溯 👍 28 👎 0

// 2022-06-19 08:13:59
//leetcode submit region begin(Prohibit modification and deletion)
func generateSentences(synonyms [][]string, text string) []string {
	// TODO

}

//leetcode submit region end(Prohibit modification and deletion)
