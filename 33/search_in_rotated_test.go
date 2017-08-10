package search

import (
	"testing"
)

func Test_seach(t *testing.T) {

	nums := []int{5, 6, 7, 1, 2, 3, 4}
	result := search(nums, 4)
	if result != 6 {
		t.Error("错误", result)
	}
}
