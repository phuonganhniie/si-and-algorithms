package leetcode_787

import "math"

type FlightInfo struct {
	nextCity, cost int
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Represent the flights information as a graph
	// Graph: [from][]{toCity, cost}
	graph := make(map[int][]FlightInfo)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]
		graph[from] = append(graph[from], FlightInfo{to, cost})
	}

	// BFS queue, each element is a tuple: {current city, cost so far, increased stop}
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{src, 0, -1}) // starting stops -1 because of since src doesn't count as a stop)

	minCostMap := make(map[int]int) // map[srcCity]{minCost}
	for i := 0; i < n; i++ {
		minCostMap[i] = math.MaxInt32
	}
	minCostMap[src] = 0

	// Time to BFS the all flights and find the minimum cost for k stops
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		currentCity, costSoFar, increasedStop := front[0], front[1], front[2]

		if increasedStop > k {
			continue
		}

		// Explore the neighbor flights in graph
		for _, info := range graph[currentCity] {
			nextCity, cost := info.nextCity, info.cost
			nextCost := cost + costSoFar
			if nextCost < minCostMap[nextCity] && increasedStop+1 <= k {
				minCostMap[nextCity] = nextCost
				queue = append(queue, [3]int{nextCity, nextCost, increasedStop + 1})
			}
		}
	}

	if minCostMap[dst] == math.MaxInt32 {
		return -1
	}
	return minCostMap[dst]
}
