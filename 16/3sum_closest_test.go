package threeSumClosest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	nums := []int{1, 2, 5, 13, -1, 3}
	result := threeSumClosest(nums, 5)
	if result != 5 {
		t.Error("test case failed")
	}
}
