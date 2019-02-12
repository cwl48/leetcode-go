package insert57

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   []Interval
		newInterval Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"test", args{[]Interval{Interval{1, 3}, Interval{6, 9}}, Interval{2, 5}}, []Interval{Interval{1, 5}, Interval{6, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
