package largestrectanglearea

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
		{"test", args{[]int{1, 1}}, 2},
		{"test", args{[]int{1}}, 1},
		{"test", args{[]int{2, 3}}, 4},
		{"test", args{[]int{}}, 0},
		{"test", args{[]int{0, 9}}, 9},
		{"test", args{[]int{2, 1, 2}}, 3},
		{"test", args{[]int{1, 2, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
