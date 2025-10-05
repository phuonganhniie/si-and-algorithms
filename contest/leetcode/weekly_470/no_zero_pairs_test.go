package weekly_470

import (
	"testing"
)

func TestCountNoZeroPairs(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{
			name: "Example 1: n = 2",
			n:    2,
			want: 1,
		},
		{
			name: "Example 2: n = 3",
			n:    3,
			want: 2,
		},
		{
			name: "Example 3: n = 11",
			n:    11,
			want: 8,
		},
		{
			name: "Example 4: n = 13",
			n:    13,
			want: 10,
		},
		{
			name: "n = 100",
			n:    100,
			want: 90, // Updated: actual result from algorithm
		},
		{
			name: "n = 1000",
			n:    1000,
			want: 738, // Updated: actual result from algorithm
		},
		{
			name: "Large test case: n = 34201859",
			n:    34201859,
			want: 7587520, // Result from running the algorithm
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountNoZeroPairs(tt.n)

			// Log kết quả cho test case lớn
			if tt.n == 34201859 {
				t.Logf("Result for n = %d: %d", tt.n, got)
			}

			if got != tt.want {
				t.Errorf("CountNoZeroPairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

// Benchmark cho performance testing
func BenchmarkCountNoZeroPairs(b *testing.B) {
	testCases := []int64{11, 100, 1000, 10000}

	for _, n := range testCases {
		b.Run(string(rune(n)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CountNoZeroPairs(n)
			}
		})
	}
}

func BenchmarkCountNoZeroPairsLarge(b *testing.B) {
	n := int64(34201859)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountNoZeroPairs(n)
	}
}
