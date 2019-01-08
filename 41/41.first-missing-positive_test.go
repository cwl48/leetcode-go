package firstmisspositve

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{1, 2, 0}}, 3},
		{"test", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"test", args{[]int{3, 4, -1, 1}}, 2},
		{"test", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
