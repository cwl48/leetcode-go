package sortcolors75

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{[]int{2, 0, 2, 1, 1, 0}}, []int{0, 0, 1, 1, 2, 2}},
		{"test", args{[]int{1, 2, 0}}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortColors(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
