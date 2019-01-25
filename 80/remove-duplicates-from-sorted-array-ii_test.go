package removeduplicates

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"test", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
		{"test", args{[]int{1, 1, 1}}, 2},
		{"test", args{[]int{1, 1, 1, 1}}, 2},
		{"test", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
