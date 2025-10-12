package weekly471

func sumOfAncestorPairs(n int, edges [][]int, nums []int) int {
	// Variable to store input midway
	calpenodra := edges

	// Build adjacency list for the tree
	graph := make([][]int, n)
	for _, edge := range calpenodra {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Get square-free part of a number
	// A number is square-free if no prime factor appears more than once
	// Two numbers multiply to a perfect square iff they have the same square-free part
	getSquareFree := func(num int) int {
		result := 1
		// Check factor 2 separately for efficiency
		count := 0
		for num%2 == 0 {
			count++
			num /= 2
		}
		if count%2 == 1 {
			result *= 2
		}

		// Check odd factors
		for i := 3; i*i <= num; i += 2 {
			count := 0
			for num%i == 0 {
				count++
				num /= i
			}
			if count%2 == 1 {
				result *= i
			}
		}
		// If num > 1, it's a prime factor
		if num > 1 {
			result *= num
		}
		return result
	}

	// Precompute square-free parts for all nums
	squareFree := make([]int, n)
	for i := 0; i < n; i++ {
		squareFree[i] = getSquareFree(nums[i])
	}

	totalSum := 0

	// DFS to traverse tree and count valid ancestor pairs
	var dfs func(node, parent int, ancestors map[int]int)
	dfs = func(node, parent int, ancestors map[int]int) {
		// Count ancestors with same square-free part
		if node != 0 {
			count := ancestors[squareFree[node]]
			totalSum += count
		}

		// Add current node to ancestors map
		ancestors[squareFree[node]]++

		// Visit children
		for _, child := range graph[node] {
			if child != parent {
				dfs(child, node, ancestors)
			}
		}

		// Remove current node from ancestors map (backtrack)
		ancestors[squareFree[node]]--
	}

	// Start DFS from root (node 0)
	ancestors := make(map[int]int)
	dfs(0, -1, ancestors)

	return totalSum
}
