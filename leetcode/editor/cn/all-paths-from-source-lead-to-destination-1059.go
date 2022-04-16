package main

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	colors := make([]int, n)
	graph := make(map[int][]int)

	for _, arr := range edges {
		s, d := arr[0], arr[1]
		graph[s] = append(graph[s], d)
	}

	var dfs func(int) bool
	dfs = func(src int) bool {
		if len(graph[src]) == 0 {
			return src == destination
		}
		colors[src] = 1
		if len(graph[src]) > 0 {
			for _, v := range graph[src] {
				if colors[v] == 1 {
					return false
				}
				if colors[v] == 0 && !dfs(v) {
					return false
				}
			}
			colors[src] = 2
		}
		return true
	}
	return dfs(source)
}
