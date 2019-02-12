package merge56

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"test", args{[]Interval{Interval{Start: 1, End: 4}, Interval{Start: 0, End: 1}}}, []Interval{Interval{Start: 0, End: 4}}},
		{"test", args{[]Interval{Interval{Start: 1, End: 3}, Interval{Start: 0, End: 4}}}, []Interval{Interval{Start: 0, End: 4}}},
		{"test", args{[]Interval{Interval{Start: 1, End: 4}, Interval{Start: 4, End: 5}}}, []Interval{Interval{Start: 1, End: 5}}},
		{"test", args{[]Interval{Interval{Start: 1, End: 4}, Interval{Start: 2, End: 3}}}, []Interval{Interval{Start: 1, End: 4}}},
		{"test", args{[]Interval{Interval{Start: 1, End: 4}, Interval{Start: 0, End: 0}}}, []Interval{Interval{Start: 1, End: 4}, Interval{Start: 0, End: 0}}},
		{"test", args{[]Interval{Interval{Start: 2, End: 3}, Interval{Start: 4, End: 5},Interval{Start: 6, End: 7},Interval{Start: 8, End: 9},
		Interval{Start: 1, End: 10}}}, []Interval{Interval{Start: 1, End: 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
