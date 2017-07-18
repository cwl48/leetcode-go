package ContainerWater

import (
	"testing"
)

func Test_maxArea(t *testing.T) {

	nums := []int{1, 2, 4, 3}
	result := maxArea(nums)
	if result != 4 {
		t.Error("test case failed")
	}

}
