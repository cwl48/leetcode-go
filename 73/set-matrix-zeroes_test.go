package setzeroes

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
		})
	}
}
