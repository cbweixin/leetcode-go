package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 你有一些长度为正整数的棍子。这些长度以数组 sticks 的形式给出， sticks[i] 是 第i个 木棍的长度。
//
// 你可以通过支付 x + y 的成本将任意两个长度为 x 和 y 的棍子连接成一个棍子。你必须连接所有的棍子，直到剩下一个棍子。
//
// 返回以这种方式将所有给定的棍子连接成一个棍子的 最小成本 。
//
//
//
// 示例 1：
//
//
// 输入：sticks = [2,4,3]
// 输出：14
// 解释：从 sticks = [2,4,3] 开始。
// 1. 连接 2 和 3 ，费用为 2 + 3 = 5 。现在 sticks = [5,4]
// 2. 连接 5 和 4 ，费用为 5 + 4 = 9 。现在 sticks = [9]
// 所有木棍已经连成一根，总费用 5 + 9 = 14
//
//
// 示例 2：
//
//
// 输入：sticks = [1,8,3,5]
// 输出：30
// 解释：从 sticks = [1,8,3,5] 开始。
// 1. 连接 1 和 3 ，费用为 1 + 3 = 4 。现在 sticks = [4,8,5]
// 2. 连接 4 和 5 ，费用为 4 + 5 = 9 。现在 sticks = [9,8]
// 3. 连接 9 和 8 ，费用为 9 + 8 = 17 。现在 sticks = [17]
// 所有木棍已经连成一根，总费用 4 + 9 + 17 = 30
//
//
// 示例 3：
//
//
// 输入：sticks = [5]
// 输出：0
// 解释：只有一根木棍，不必再连接。总费用 0
//
//
//
//
// 提示：
//
//
// 1 <= sticks.length <= 10⁴
// 1 <= sticks[i] <= 10⁴
//
// Related Topics 贪心 数组 堆（优先队列） 👍 60 👎 0

// 2022-05-28 07:57:37
// leetcode submit region begin(Prohibit modification and deletion)
//  sort
func connectSticks2(sticks []int) int {
	sort.Ints(sticks)
	stack := make([]int, 0)

	pop := func(arr *[]int) int {
		res := (*arr)[0]
		*arr = (*arr)[1:]
		return res
	}

	leftPop := func(s1, s2 *[]int) int {
		var res int
		if len(*s1) == 0 {
			res = pop(s2)
		} else if len(*s2) == 0 {
			res = pop(s1)
		} else if (*s1)[0] < (*s2)[0] {
			res = pop(s1)
		} else {
			res = pop(s2)
		}
		return res
	}

	res := 0
	for len(sticks)+len(stack) > 1 {
		a := 0
		a += leftPop(&stack, &sticks)
		a += leftPop(&stack, &sticks)
		res += a
		stack = append(stack, a)
	}
	return res

}

func connectSticks(sticks []int) int {
	h := IntHeap(sticks)
	heap.Init(&h)
	res := 0
	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		res += a + b
		heap.Push(&h, a+b)
	}
	return res
}

// knowledge how to implement heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	sticks := []int{2, 4, 3}
	res := connectSticks(sticks)
	fmt.Println(res)
}
