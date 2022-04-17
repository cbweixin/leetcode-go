package main

type Adj struct {
	ajs []int
}

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

func leadsToDestination1(n int, edges [][]int, source int, destination int) bool {
	colors := make([]int, n)
	// notice here I can not do make(map[int] Adj), have to make it a pointer type
	// https://stackoverflow.com/questions/40578646/golang-i-have-a-map-of-int-to-struct-why-cant-i-directly-modify-a-field-in-a
	// You are storing a struct by value which means that accession of that struct in the map gives you a copy of the value. This is why when you modify it, the struct in the map remains unmutated until you overwrite it with the new copy.
	//
	// As RickyA pointed out in the comment, you can store the pointer to the struct instead and this allows direct modification of the struct being referenced by the stored struct pointer.
	//
	// i.e. map[whatever]*struct instead of map[whatever]struct
	graph := make(map[int]*Adj)

	for i := 0; i < n; i++ {
		graph[i] = &Adj{ajs: []int{}}
	}

	for _, arr := range edges {
		s, d := arr[0], arr[1]
		graph[s].ajs = append(graph[s].ajs, d)
	}

	var dfs func(int) bool
	dfs = func(src int) bool {
		if len(graph[src].ajs) == 0 {
			return src == destination
		}
		colors[src] = 1
		if len(graph[src].ajs) > 0 {
			for _, v := range graph[src].ajs {
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
