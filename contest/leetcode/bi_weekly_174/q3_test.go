package biweekly174

import "testing"

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target1  int
		target2  int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 4},
			target1:  1,
			target2:  5,
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 0, 0},
			target1:  1,
			target2:  0,
			expected: 3,
		},
		{
			name:     "Example 3 - No valid partition",
			nums:     []int{7},
			target1:  1,
			target2:  7,
			expected: 0,
		},
		{
			name:     "Single element matching target1",
			nums:     []int{5},
			target1:  5,
			target2:  3,
			expected: 1,
		},
		{
			name:     "Two elements, two blocks",
			nums:     []int{3, 5},
			target1:  3,
			target2:  5,
			expected: 1,
		},
		{
			name:     "All zeros with target1=0",
			nums:     []int{0, 0, 0},
			target1:  0,
			target2:  1,
			expected: 1, // Only [0,0,0] with XOR=0=target1 is valid; can't get target2=1 from zeros
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countPartitions(tt.nums, tt.target1, tt.target2)
			if result != tt.expected {
				t.Errorf("countPartitions() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
