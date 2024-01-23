package bfs

/** Pseudo Code
function rottenOranges(grid):
    if grid is empty:
        return -1

    rows = number of rows in grid
    cols = number of columns in grid
    freshCount = 0
    time = 0
    queue = empty queue

    // Initialize the queue with all initially rotten oranges
    // and count the fresh oranges
    for each row in grid:
        for each col in grid:
            if grid[row][col] == 2:
                add (row, col) to queue
            else if grid[row][col] == 1:
                freshCount++

    // Directions array to check 4-directionally adjacent cells
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

    // BFS iteration
    while queue is not empty and freshCount > 0:
        size = size of queue
        for i from 0 to size:
            (x, y) = dequeue from queue

            // Check all 4-directionally adjacent cells
            for each (dx, dy) in directions:
                newX = x + dx
                newY = y + dy

                // Check if the new cell is within the grid and is a fresh orange
                if 0 <= newX < rows and 0 <= newY < cols and grid[newX][newY] == 1:
                    grid[newX][newY] = 2  // Make the orange rotten
                    enqueue (newX, newY) into queue
                    freshCount--

        // Increment time after each minute of processing
        time++

    // If there are still fresh oranges, return -1
    if freshCount > 0:
        return -1

    return time
*/

type position struct {
	x, y int
}

func OrangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	rows := len(grid)
	cols := len(grid[0])
	freshOranges := 0
	time := 0
	queue := make([]position, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, position{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	directions := []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // up, right, down, left

	for len(queue) > 0 && freshOranges > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			currentDir := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				newXDir := currentDir.x + dir.x
				newYDir := currentDir.y + dir.y

				if newXDir >= 0 && newXDir < rows && newYDir >= 0 && newYDir < cols && grid[newXDir][newYDir] == 1 {
					grid[newXDir][newYDir] = 2
					queue = append(queue, position{newXDir, newYDir})
					freshOranges--
				}
			}
		}
		time++
	}

	if freshOranges > 0 {
		return -1
	}
	return time
}
