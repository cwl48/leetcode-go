package removeDuplicates

import "testing"

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 5, 6, 6, 6}
	lens := removeDuplicates(nums)

	if lens != 3 {
		t.Error("错误", lens)
	}
}
