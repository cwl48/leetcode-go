package maxprofit3

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"test", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"test", args{[]int{3, 3, 5, 0, 0, 3, 1, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
