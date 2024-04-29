package fiverr

import "testing"

func TestGetCableLength(t *testing.T) {
	tests := []struct {
		rowInput   int
		colInput   int
		imageInput []string
		want       int
	}{
		{2, 3, []string{"ABC", "CBC"}, 14},
		{4, 4, []string{"AABC", "CAAC", "BCCA", "CABA"}, 228},
	}

	for _, tt := range tests {
		got := GetCableLength(tt.rowInput, tt.colInput, tt.imageInput)
		if got != tt.want {
			t.Errorf("GetCableLength(%v, %v, %v) failed, want %d - got %d", tt.rowInput, tt.colInput, tt.imageInput, tt.want, got)
		}
	}
}
