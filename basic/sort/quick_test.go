package sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				arr:   []int{1, 3, 5, 7, 2, 3, 8, 9, 12, 34, 8, 0, 2, 55, 2, 3, 3, 6},
				left:  0,
				right: 17,
			},
			want: []int{0, 1, 2, 2, 2, 3, 3, 3, 3, 5, 6, 7, 8, 8, 9, 12, 34, 55},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
