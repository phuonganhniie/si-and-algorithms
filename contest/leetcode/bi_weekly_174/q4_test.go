package biweekly174

import (
	"reflect"
	"testing"
)

func TestTreeTransform(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		start    string
		target   string
		expected []int
	}{
		{
			name:     "Example 1",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			start:    "010",
			target:   "100",
			expected: []int{0},
		},
		{
			name:     "Example 2",
			n:        7,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {3, 5}, {1, 6}},
			start:    "0011000",
			target:   "0010001",
			expected: []int{1, 2, 5},
		},
		{
			name:     "Example 3 - Impossible",
			n:        2,
			edges:    [][]int{{0, 1}},
			start:    "00",
			target:   "01",
			expected: []int{-1},
		},
		{
			name:     "Already equal",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			start:    "101",
			target:   "101",
			expected: []int{},
		},
		{
			name:     "Single edge both flip",
			n:        2,
			edges:    [][]int{{0, 1}},
			start:    "00",
			target:   "11",
			expected: []int{0},
		},
		{
			name:     "Star graph",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			start:    "0111",
			target:   "0000",
			expected: []int{-1}, // need to flip 3 nodes, but each edge flips 2
		},
		{
			name:   "Star graph - possible",
			n:      4,
			edges:  [][]int{{0, 1}, {0, 2}, {0, 3}},
			start:  "0110",
			target: "1001",
			// diff = [1, 1, 1, 1], flip edges 0,1,2,3? No, we have 3 edges only
			// Actually: flip edge0 (0,1), edge1 (0,2), edge2 (0,3)
			// Node 0 flipped 3 times = 1 flip net, nodes 1,2,3 each flipped once
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := treeTransform(tt.n, tt.edges, tt.start, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("treeTransform() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
