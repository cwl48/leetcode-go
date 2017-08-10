package twoSum

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	arr := twoSum(nums, 9)

	if arr[0] != 0 || arr[1] != 1 {
		t.Error("twoSum error")
	}
}
