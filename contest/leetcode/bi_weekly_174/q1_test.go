package biweekly174

import (
	"reflect"
	"testing"
)

func TestBestTower(t *testing.T) {
	tests := []struct {
		name     string
		towers   [][]int
		center   []int
		radius   int
		expected []int
	}{
		{
			name:     "Example 1 - All reachable, max quality wins",
			towers:   [][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}},
			center:   []int{1, 1},
			radius:   2,
			expected: []int{3, 1},
		},
		{
			name:     "Example 2 - Tie, lexicographically smaller wins",
			towers:   [][]int{{1, 3, 4}, {2, 2, 4}, {4, 4, 7}},
			center:   []int{0, 0},
			radius:   5,
			expected: []int{1, 3},
		},
		{
			name:     "Example 3 - No tower reachable",
			towers:   [][]int{{5, 6, 8}, {0, 3, 5}},
			center:   []int{1, 2},
			radius:   1,
			expected: []int{-1, -1},
		},
		{
			name:     "Single tower reachable",
			towers:   [][]int{{0, 0, 10}},
			center:   []int{0, 0},
			radius:   0,
			expected: []int{0, 0},
		},
		{
			name:     "Tie with same x, smaller y wins",
			towers:   [][]int{{1, 2, 5}, {1, 1, 5}},
			center:   []int{0, 0},
			radius:   10,
			expected: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bestTower(tt.towers, tt.center, tt.radius)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("bestTower() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
