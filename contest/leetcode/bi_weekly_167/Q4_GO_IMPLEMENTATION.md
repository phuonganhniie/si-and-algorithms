# Q4 Go Implementation Update

## Changes Made

### Updated `q4.go`

Replaced the brute force O(2^n) approach with the optimized **Binary Search + Bipartite Graph** algorithm (O(n² log D)) to match the Python implementation.

**Key Components:**

1. **`maxPartitionFactor(points [][]int) int`**
   - Main function implementing binary search on distance values
   - Collects all unique Manhattan distances
   - Binary searches to find maximum achievable partition factor

2. **`canAchieve(points [][]int, minFactor int) bool`**
   - Checks if a given partition factor is achievable
   - Builds conflict graph where edges connect points with distance < minFactor
   - Uses BFS to check if graph is bipartite (2-colorable)
   - Handles edge case: when no conflicts exist, manually assigns one node to different color

3. **Helper Functions:**
   - `sortInts(arr []int)`: Simple bubble sort for sorting distances
   - `abs(x int) int`: Absolute value function

### Algorithm Flow

```
1. Calculate all pairwise Manhattan distances: O(n²)
2. Get unique distances and sort them: O(n² log n²)
3. Binary search on distance array:
   - For each candidate distance:
     - Build conflict graph (edges where distance < candidate)
     - Check if graph is bipartite using BFS
     - If bipartite AND both groups non-empty: try larger distance
     - Otherwise: try smaller distance
4. Return maximum achievable partition factor
```

### Complexity Analysis

- **Time Complexity:** O(n² log D)
  - n: number of points
  - D: number of unique distances (at most n²)
  - O(n²) to compute distances
  - O(log D) binary search iterations
  - O(n²) per iteration for graph building and BFS

- **Space Complexity:** O(n²)
  - Conflict graph storage

### Test Results

All tests passing:

```
✅ Example 1: [[0,0],[0,2],[2,0],[2,2]] → 4
✅ Example 2: [[0,0],[0,1],[10,0]] → 11
✅ Edge case n=2: [[0,0],[1,1]] → 0
✅ Three points in line: [[0,0],[1,0],[2,0]] → 2
✅ Four points square: [[0,0],[0,1],[1,0],[1,1]] → 2
✅ Large case 1 (n=16): → 23100
✅ Large case 2 (n=16): → 39190
```

### Updated Test File `q4_test.go`

Enhanced with:
- Table-driven tests for better organization
- Large test cases (n=16) matching Python version
- Unit tests for `canAchieve` helper function
- Clear test names and descriptions

## Comparison: Go vs Python

Both implementations now use the same algorithm:

| Aspect | Go | Python |
|--------|----|----|
| Algorithm | Binary Search + Bipartite | Binary Search + Bipartite |
| Time Complexity | O(n² log D) | O(n² log D) |
| Space Complexity | O(n²) | O(n²) |
| Sorting | Custom bubble sort | Built-in sorted() |
| BFS | Manual queue with slice | collections.deque |
| Graph | [][]int adjacency list | list of lists |

## Why This Approach Works

**Bipartite Graph Theorem:**
- A graph is bipartite ⟺ it can be 2-colored with no adjacent vertices having the same color
- In our problem: "adjacent" means Manhattan distance < k
- If we can 2-color the conflict graph → we can partition points into two groups where all intra-group distances ≥ k

**Binary Search Optimization:**
- Monotonic property: if partition factor k is achievable, all factors < k are also achievable
- Instead of trying all O(n²) possible factors, we binary search over O(log n²) candidates
- Each check takes O(n²), giving total O(n² log n)

**Edge Case Handling:**
- When conflict graph has no edges (all distances ≥ candidate), all nodes receive color 0 during BFS
- We detect this and manually assign one node to color 1
- Ensures both groups are non-empty (required by problem)

## Running Tests

```bash
# Run all tests
cd d:\happycoding\si-and-algorithms\contest\leetcode\bi_weekly_167
go test -v

# Run specific test
go test -v -run TestMaxPartitionFactor

# Run with coverage
go test -cover
```

## Files Modified

1. `q4.go` - Complete rewrite with optimized algorithm
2. `q4_test.go` - Enhanced test suite with large cases
3. Documentation updated in `README.md` and `README_VI.md`
