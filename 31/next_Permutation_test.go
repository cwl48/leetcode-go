package nextpermutation

import "testing"

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2}

	result := nextPermutation(nums)
	if result[0] != 2 || result[1] != 1 {
		t.Error("错误", result)
	}
}
