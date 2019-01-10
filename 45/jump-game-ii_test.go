package jumpgameli

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{2, 3, 1, 1, 4}}, 2},
		{"test", args{[]int{2, 1, 1, 1, 4}}, 3},
		{"test", args{[]int{1, 2}}, 1},
		{"test", args{[]int{1, 2, 3}}, 2},
		{"test", args{[]int{3, 2, 1}}, 1},
		{"test", args{[]int{1, 1, 1, 2, 1}}, 4},
		{"test", args{[]int{1, 1, 2, 1, 1}}, 3},
		{"test", args{[]int{1, 2, 0, 1}}, 2},
		{"test", args{[]int{2, 1, 1, 1, 1}}, 3},
		{"test", args{[]int{2, 0, 2, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
