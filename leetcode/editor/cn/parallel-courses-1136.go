package main

type Node struct {
	Idx  int
	time int
}

func minimumSemesters(n int, relations [][]int) int {
	inDeg, outDeg := make(map[int]map[int]bool), make(map[int]map[int]bool)

	for _, v := range relations {
		s, t := v[0], v[1]
		if _, ok := inDeg[t]; !ok {
			inDeg[t] = make(map[int]bool)
		}
		if _, ok := outDeg[s]; !ok {
			outDeg[s] = make(map[int]bool)
		}
		inDeg[t][s] = true
		outDeg[s][t] = true
	}

	que := make([]Node, 0)
	seen := make(map[int]bool)

	for i := 1; i <= n; i++ {
		if _, ok := inDeg[i]; !ok {
			que = append(que, Node{Idx: i, time: 1})
		}
	}

	res := 0
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		res = node.time
		if _, ok := seen[node.Idx]; !ok {
			seen[node.Idx] = true
			for k, _ := range outDeg[node.Idx] {
				delete(inDeg[k], node.Idx)
				if len(inDeg[k]) == 0 {
					delete(inDeg, k)
					que = append(que, Node{Idx: k, time: node.time + 1})
				}
			}
			delete(outDeg, node.Idx)
		} else {
			return -1
		}

	}

	if len(seen) == n {
		return res
	} else {
		return -1
	}
}

func main() {
	n := 3
	relations := [][]int{{1, 3}, {2, 3}}
	res := minimumSemesters(n, relations)
	println(res)

	n = 3
	relations = [][]int{{1, 2}, {2, 3}, {3, 1}}
	res = minimumSemesters(n, relations)
	println(res)
}
