package findpairs

import (
	"testing"
)

func Test_findPairs(t *testing.T) {
	nums := []int{1, 3, 1, 5, 4}
	r := findPairs(nums, 2)
	if r != 2 {
		t.Error("case not pass")
	}
}
