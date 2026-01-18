package biweekly174

import "sort"

func treeTransform(n int, edges [][]int, start string, target string) []int {
	// Calculate which nodes need to be flipped
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		if start[i] != target[i] {
			diff[i] = 1
		}
	}

	// Build adjacency list with edge indices
	type edge struct {
		to    int
		index int
	}
	adj := make([][]edge, n)
	for i, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], edge{v, i})
		adj[v] = append(adj[v], edge{u, i})
	}

	// BFS to find parent and edge to parent for each node
	parent := make([]int, n)
	edgeToParent := make([]int, n)
	parent[0] = -1
	edgeToParent[0] = -1

	order := make([]int, 0, n) // BFS order
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, e := range adj[node] {
			if !visited[e.to] {
				visited[e.to] = true
				parent[e.to] = node
				edgeToParent[e.to] = e.index
				queue = append(queue, e.to)
			}
		}
	}

	// Process nodes in reverse BFS order (leaves first)
	result := []int{}
	for i := n - 1; i >= 1; i-- {
		node := order[i]
		if diff[node] == 1 {
			// Must toggle edge to parent
			edgeIdx := edgeToParent[node]
			result = append(result, edgeIdx)
			diff[node] ^= 1
			diff[parent[node]] ^= 1
		}
	}

	// Check if root is satisfied
	if diff[0] == 1 {
		return []int{-1}
	}

	// Sort result in increasing order
	sort.Ints(result)
	return result
}
