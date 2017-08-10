package SumOfTwoIntegers

import (
	"testing"
)

func Test_getSum(t *testing.T) {
	a := 2
	b := 6
	c := getSum(a, b)
	if c != 8 {
		t.Error("错误", c)
	}
}
