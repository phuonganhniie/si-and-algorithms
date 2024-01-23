package bfs

func SimpleBFS(u int, graph map[int][]int) []int {
	rs := []int{}
	queue := []int{}
	visited := make(map[int]bool)

	queue = append(queue, u)
	visited[u] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		rs = append(rs, node)

		for _, v := range graph[node] {
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}

	return rs
}
