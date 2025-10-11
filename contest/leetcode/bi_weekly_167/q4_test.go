package biweekly167

import "testing"

func TestMaxPartitionFactor(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}},
			expected: 4,
		},
		{
			name:     "Example 2",
			points:   [][]int{{0, 0}, {0, 1}, {10, 0}},
			expected: 11,
		},
		{
			name:     "Edge case: n = 2",
			points:   [][]int{{0, 0}, {1, 1}},
			expected: 0,
		},
		{
			name:     "Three points in a line",
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}},
			expected: 2,
		},
		{
			name:     "Four points forming a square",
			points:   [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxPartitionFactor(tt.points)
			if result != tt.expected {
				t.Errorf("maxPartitionFactor(%v) = %d, want %d", tt.points, result, tt.expected)
			}
		})
	}
}

func TestMaxPartitionFactorLargeCases(t *testing.T) {
	t.Run("Large case 1 (n=16)", func(t *testing.T) {
		points := [][]int{
			{3874, -65073}, {17744, -88825}, {51615, -6352}, {64357, -9827},
			{97388, -493}, {7489, -47612}, {-63730, 11026}, {-52374, -55516},
			{-6310, 90173}, {22797, -39450}, {-22314, -94950}, {-15225, 79748},
			{69933, -32347}, {41440, -9644}, {62951, -58618}, {72200, 23010},
		}
		result := maxPartitionFactor(points)
		t.Logf("Large case 1 (n=16) result: %d", result)
		if result == 0 {
			t.Errorf("Large case 1: got 0, expected non-zero result")
		}
	})

	t.Run("Large case 2 (n=16)", func(t *testing.T) {
		points := [][]int{
			{-85401, -17368}, {46585, -69810}, {-36644, -97571}, {-29903, -81002},
			{-41885, 70137}, {-96898, -36499}, {-92734, -61360}, {-954, 36272},
			{39099, 36602}, {1986, 58859}, {97516, 22190}, {-56062, 22948},
			{-50995, 12805}, {-13759, 35414}, {36630, 18091}, {83194, 78635},
		}
		result := maxPartitionFactor(points)
		t.Logf("Large case 2 (n=16) result: %d", result)
		if result == 0 {
			t.Errorf("Large case 2: got 0, expected non-zero result")
		}
	})
}

func TestCanAchieve(t *testing.T) {
	t.Run("Simple bipartite case", func(t *testing.T) {
		points := [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
		result := canAchieve(points, 4)
		if !result {
			t.Errorf("canAchieve for partition factor 4 should be true")
		}
	})

	t.Run("Can achieve with all conflicts", func(t *testing.T) {
		points := [][]int{{0, 0}, {0, 1}, {1, 0}}
		// All distances: 1, 1, 2 (all < 10)
		// With minFactor=10, all pairs have edges (conflict if dist < 10)
		// Graph: 0-1, 0-2, 1-2 (complete triangle)
		// Not bipartite, should return false
		result := canAchieve(points, 10)
		if result {
			t.Errorf("canAchieve for partition factor 10 should be false (complete triangle, not bipartite)")
		}
	})

	t.Run("Can achieve small factor", func(t *testing.T) {
		points := [][]int{{0, 0}, {1, 0}, {2, 0}}
		result := canAchieve(points, 1)
		if !result {
			t.Errorf("canAchieve for partition factor 1 should be true")
		}
	})
}
