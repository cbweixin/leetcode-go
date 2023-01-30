package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 2380 👎 0

//2023-01-30 10:52:49

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1 -> 2 -> 3 -> 4 -> 5, n= 2, remove 4
	// l         r
	//           l    ()    r

	// l, r are reference
	l, r := head, head
	for step := 0; step < n; step++ {
		r = r.Next
	}
	if r == nil {
		return l.Next
	}
	for r.Next != nil {
		l, r = l.Next, r.Next
	}
	t := l.Next
	l.Next = t.Next
	t.Next = nil

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
