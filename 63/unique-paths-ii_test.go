package uniquepaths2

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"test", args{[][]int{{0}, {0}}}, 1},
		{"test", args{[][]int{{0, 0}}}, 1},
		{"test", args{[][]int{{0, 0}, {1, 1}, {0, 0}}}, 0},
		{"test", args{[][]int{{1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
