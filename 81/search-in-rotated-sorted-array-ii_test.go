package search81

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {"test", args{[]int{2, 5, 6, 0, 0, 1, 2}, 0}, true},
		// {"test", args{[]int{2, 5, 6, 0, 0, 1, 2}, 3}, false},
		// {"test", args{[]int{1, 3}, 3}, true},
		// {"test", args{[]int{1, 1, 3, 1}, 3}, true},
		{"test", args{[]int{1, 3, 1, 1, 1}, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
