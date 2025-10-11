package biweekly167

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	if n == 2 {
		return 0
	}

	distanceMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			distanceMap[dist] = true
		}
	}

	distances := make([]int, 0, len(distanceMap))
	for dist := range distanceMap {
		distances = append(distances, dist)
	}
	sortInts(distances)

	left, right := 0, len(distances)-1
	result := 0

	for left <= right {
		mid := (left + right) / 2
		candidate := distances[mid]

		if canAchieve(points, candidate) {
			result = candidate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func canAchieve(points [][]int, minFactor int) bool {
	n := len(points)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			if dist < minFactor {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	bfs := func(start int) bool {
		queue := []int{start}
		color[start] = 0

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if color[neighbor] == -1 {
					color[neighbor] = 1 - color[node]
					queue = append(queue, neighbor)
				} else if color[neighbor] == color[node] {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if color[i] == -1 {
			if !bfs(i) {
				return false
			}
		}
	}

	hasEdge := false
	for i := 0; i < n; i++ {
		if len(graph[i]) > 0 {
			hasEdge = true
			break
		}
	}
	if !hasEdge {
		color[0] = 1
	}

	count0, count1 := 0, 0
	for _, c := range color {
		if c == 0 {
			count0++
		} else if c == 1 {
			count1++
		}
	}

	return count0 > 0 && count1 > 0
}

func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
