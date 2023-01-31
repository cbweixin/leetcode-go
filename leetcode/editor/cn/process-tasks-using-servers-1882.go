package main

import (
	"container/heap"
	"fmt"
	"math"
)

//给你两个 下标从 0 开始 的整数数组 servers 和 tasks ，长度分别为 n 和 m 。servers[i] 是第 i 台服务器的 权重 ，而
//tasks[j] 是处理第 j 项任务 所需要的时间（单位：秒）。
//
// 你正在运行一个仿真系统，在处理完所有任务后，该系统将会关闭。每台服务器只能同时处理一项任务。第 0 项任务在第 0 秒可以开始处理，相应地，第 j 项任务
//在第 j 秒可以开始处理。处理第 j 项任务时，你需要为它分配一台 权重最小 的空闲服务器。如果存在多台相同权重的空闲服务器，请选择 下标最小 的服务器。如果一
//台空闲服务器在第 t 秒分配到第 j 项任务，那么在 t + tasks[j] 时它将恢复空闲状态。
//
// 如果没有空闲服务器，则必须等待，直到出现一台空闲服务器，并 尽可能早 地处理剩余任务。 如果有多项任务等待分配，则按照 下标递增 的顺序完成分配。
//
// 如果同一时刻存在多台空闲服务器，可以同时将多项任务分别分配给它们。
//
// 构建长度为 m 的答案数组 ans ，其中 ans[j] 是第 j 项任务分配的服务器的下标。
//
// 返回答案数组 ans 。
//
//
//
// 示例 1：
//
//
//输入：servers = [3,3,2], tasks = [1,2,3,2,1,2]
//输出：[2,2,0,2,1,2]
//解释：事件按时间顺序如下：
//- 0 秒时，第 0 项任务加入到任务队列，使用第 2 台服务器处理到 1 秒。
//- 1 秒时，第 2 台服务器空闲，第 1 项任务加入到任务队列，使用第 2 台服务器处理到 3 秒。
//- 2 秒时，第 2 项任务加入到任务队列，使用第 0 台服务器处理到 5 秒。
//- 3 秒时，第 2 台服务器空闲，第 3 项任务加入到任务队列，使用第 2 台服务器处理到 5 秒。
//- 4 秒时，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 5 秒。
//- 5 秒时，所有服务器都空闲，第 5 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
//
// 示例 2：
//
//
//输入：servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
//输出：[1,4,1,4,1,3,2]
//解释：事件按时间顺序如下：
//- 0 秒时，第 0 项任务加入到任务队列，使用第 1 台服务器处理到 2 秒。
//- 1 秒时，第 1 项任务加入到任务队列，使用第 4 台服务器处理到 2 秒。
//- 2 秒时，第 1 台和第 4 台服务器空闲，第 2 项任务加入到任务队列，使用第 1 台服务器处理到 4 秒。
//- 3 秒时，第 3 项任务加入到任务队列，使用第 4 台服务器处理到 7 秒。
//- 4 秒时，第 1 台服务器空闲，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 9 秒。
//- 5 秒时，第 5 项任务加入到任务队列，使用第 3 台服务器处理到 7 秒。
//- 6 秒时，第 6 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
//
//
//
// 提示：
//
//
// servers.length == n
// tasks.length == m
// 1 <= n, m <= 2 * 10⁵
// 1 <= servers[i], tasks[j] <= 2 * 10⁵
//
//
// Related Topics 数组 堆（优先队列） 👍 61 👎 0

//2023-01-31 12:24:00

// leetcode submit region begin(Prohibit modification and deletion)
func assignTasks1(servers []int, tasks []int) []int {
	// heap
	release := make(map[int][]server, 0)
	availble, curTime, i := make(ServerHeap, 0), 0, 0
	for k, v := range servers {
		availble.Push(server{weight: v, index: k})
	}
	heap.Init(&availble)
	res := make([]int, len(tasks))

	minKey := func(x map[int][]server) int {
		res := math.MaxInt
		for k := range x {
			if k < res {
				res = k
			}
		}
		return res
	}

	for i < len(tasks) {
		if v, ok := release[curTime]; ok {
			for _, s := range v {
				heap.Push(&availble, s)
			}
			// don't forget delete it
			delete(release, curTime)
		}

		if len(availble) == 0 {
			curTime = minKey(release)
			continue
		}

		for i <= curTime && len(availble) > 0 {
			s := heap.Pop(&availble).(server)
			res[i] = s.index
			t := curTime + tasks[i]
			if v, ok := release[t]; ok {
				release[t] = append(v, server{weight: s.weight, index: s.index})
			} else {
				release[t] = []server{{weight: s.weight, index: s.index}}
			}
			i += 1
			if i == len(tasks) {
				break
			}

		}
		curTime += 1
	}
	return res
}

func assignTasks(servers []int, tasks []int) []int {
	free, busy, curTime, l, taskIdx := make(ServerHeap, 0), make(ServerHeap, 0), 0, len(tasks), 0
	res := make([]int, l)

	for i, w := range servers {
		free.Push(server{weight: w, index: i})
	}

	heap.Init(&free)
	heap.Init(&busy)

	for taskIdx < l {
		for len(busy) > 0 && busy[0].weight <= curTime {
			s := heap.Pop(&busy).(server)
			heap.Push(&free, server{weight: servers[s.index], index: s.index})
		}

		if len(free) == 0 {
			curTime = busy[0].weight
			continue
		}

		for taskIdx <= curTime && len(free) > 0 {
			s := heap.Pop(&free).(server)
			res[taskIdx] = s.index
			heap.Push(&busy, server{weight: curTime + tasks[taskIdx], index: s.index})
			taskIdx += 1
			if taskIdx == l {
				break
			}
		}

		curTime += 1
	}

	return res

}

type server struct {
	weight int
	index  int
}

type ServerHeap []server

func (h ServerHeap) Len() int {
	return len(h)
}

func (h ServerHeap) Less(i, j int) bool {
	// bugfixed , if 2 server has same weight, but different index, without this fix , it won't work
	// [61, 8], [61, 2], but I hope [61,2], [61,8]
	if h[i].weight == h[j].weight {
		return h[i].index < h[j].index
	}
	return h[i].weight < h[j].weight
}

func (h ServerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ServerHeap) Push(x interface{}) {
	*h = append(*h, x.(server))
}

func (h *ServerHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(assignTasks(
		[]int{74, 57, 61, 82, 67, 97, 67, 21, 61, 79, 21, 50, 14, 88, 48, 52, 76, 64},
		[]int{21, 100, 48, 64, 20, 8, 28, 10, 3, 63, 7},
	),
	)
}
