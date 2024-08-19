package queue

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestReverseKElements(t *testing.T) {
	tests := []struct {
		input *helper.Queue
		k     int
		want  *helper.Queue
	}{
		{
			&helper.Queue{
				Size:  4,
				Data:  [100]interface{}{1, 2, 3, 4},
				Front: 0,
				Back:  0,
			},
			2,
			&helper.Queue{
				Size:  4,
				Data:  [100]interface{}{2, 1, 3, 4},
				Front: 0,
				Back:  4},
		},
		{
			&helper.Queue{
				Size:  5,
				Data:  [100]interface{}{4, 8, 2, 7, 1},
				Front: 0,
				Back:  0,
			},
			3,
			&helper.Queue{
				Size:  5,
				Data:  [100]interface{}{2, 8, 4, 7, 1},
				Front: 0,
				Back:  5,
			},
		},
		{
			&helper.Queue{
				Size:  5,
				Data:  [100]interface{}{1, 2, 3, 4, 5},
				Front: 0,
				Back:  0,
			},
			5,
			&helper.Queue{
				Size:  5,
				Data:  [100]interface{}{5, 4, 3, 2, 1},
				Front: 0,
				Back:  5,
			},
		},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseKElementInQueue(tt.input, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseKElementInQueue test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
