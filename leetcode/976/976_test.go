package leetcode_976

import "testing"

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Smallest valid triangle",
			input:    []int{2, 3, 4}, // 2 + 3 > 4 => 9
			expected: 9,
		},
		{
			name:     "No triangle possible (sorted descending)",
			input:    []int{5, 4, 10}, // 4 + 5 !> 10
			expected: 0,
		},
		{
			name:     "No triangle possible (all same)",
			input:    []int{1, 1, 2}, // 1 + 1 !> 2
			expected: 0,
		},
		{
			name:     "Multiple triangles choose largest perimeter",
			input:    []int{2, 1, 2, 4, 3}, // triangles: (2,3,4)=9, (1,2,2)=5
			expected: 9,
		},
		{
			name:     "Large numbers",
			input:    []int{3000000, 3000000, 3000000, 2999999}, // pick three largest equal => 9000000
			expected: 9000000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := largestPerimeter(tt.input)
			if got != tt.expected {
				t.Errorf("largestPerimeter(%v) = %d, expected %d", tt.input, got, tt.expected)
			}
		})
	}
}
