package removeElement

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{1, 2, 1, 3}
	result := removeElement(nums, 2)
	if result != 3 {
		t.Error("错误", result)
	}
}
