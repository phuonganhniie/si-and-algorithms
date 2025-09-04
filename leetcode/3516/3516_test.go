package leetcode_3516

import "testing"

func TestFindClosest(t *testing.T) {
	tests := []struct{
		name string
		x, y, z int
		expected int
	}{
		{
			name: "Example 1: Person 1 reach first",
			x: 2,
			y: 7,
			z: 4,
			expected: 1,
		},
		{
			name: "Example 2: Person 2 reach first",
			x: 2,
			y: 5,
			z: 6,
			expected: 2,
		},
		{
			name: "Example 3: Both reach person 3",
			x: 1,
			y: 5,
			z: 3,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findClosest(tt.x, tt.y, tt.z)
			if result != tt.expected {
				t.Errorf("findClosest() = %d, want %d", result, tt.expected)
			}
		})
	}
}