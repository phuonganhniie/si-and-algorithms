package leetcode_207

func CanFinish(numCourses int, prerequisites [][]int) bool {
	adList := make([][]int, numCourses)
	for _, pre := range prerequisites {
		course := pre[0]
		prereq := pre[1]
		adList[course] = append(adList[course], prereq)
	}

	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if dfs(adList, visited, i) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, visited []int, course int) bool {
	if visited[course] == -1 {
		return true
	}
	if visited[course] == 1 {
		return false
	}

	visited[course] = -1
	for _, pre := range graph[course] {
		if dfs(graph, visited, pre) {
			return true
		}
	}
	visited[course] = 1
	return false
}
