package leetcode

/*
Description: https://leetcode.com/contest/weekly-contest-410/problems/count-the-number-of-good-nodes/description/

There is an undirected tree with n nodes labeled from 0 to n - 1, and rooted at node 0.
You are given a 2D integer array edges of length n - 1, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

A node is good if all the subtrees rooted at its children have the same size.

Return the number of good nodes in the given tree.

A subtree of treeName is a tree consisting of a node in treeName and all of its descendants.

Example 1:
Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
Output: 7

Example 2:
Input: edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]
Output: 6

Example 3:
Input: edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]
Output: 12
*/

func countGoodNodesGPT(edges [][]int) int {
	n := len(edges) + 1
	tree := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}

	subtreeSize := make([]int, n)
	visited := make([]bool, n)

	var dfs func(node int) int
	dfs = func(node int) int {
		visited[node] = true
		size := 1
		for _, neighbor := range tree[node] {
			if !visited[neighbor] {
				size += dfs(neighbor)
			}
		}
		subtreeSize[node] = size
		return size
	}

	dfs(0)

	var isGoodNode func(node int) bool
	isGoodNode = func(node int) bool {
		if len(tree[node]) == 1 && node != 0 {
			return true
		}
		childSizes := make(map[int]int)
		for _, neighbor := range tree[node] {
			if subtreeSize[neighbor] < subtreeSize[node] {
				childSizes[subtreeSize[neighbor]]++
			}
		}
		return len(childSizes) <= 1
	}

	count := 0
	for i := 0; i < n; i++ {
		if isGoodNode(i) {
			count++
		}
	}

	return count
}

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	list := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		list[u] = append(list[u], v)
		list[v] = append(list[v], u)
	}

	subtreeLen := make([]int, n)
	visited := make([]bool, n)

	dfs(0, visited, list, subtreeLen)

	var isGoodNode func(node int) bool
	isGoodNode = func(node int) bool {
		if len(list[node]) == 1 && node != 0 {
			return true
		}
		childMap := make(map[int]int)
		for _, neighbor := range list[node] {
			if subtreeLen[neighbor] < subtreeLen[node] {
				childMap[subtreeLen[neighbor]]++
			}
		}
		return len(childMap) <= 1
	}

	count := 0
	for i := 0; i < n; i++ {
		if isGoodNode(i) {
			count++
		}
	}
	return count
}

func dfs(node int, visited []bool, list [][]int, subtreeLen []int) int {
	visited[node] = true
	size := 1
	for _, neighbor := range list[node] {
		if !visited[neighbor] {
			size += dfs(neighbor, visited, list, subtreeLen)
		}
	}
	subtreeLen[node] = size
	return size
}
