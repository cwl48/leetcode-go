package medianof

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {

	nums1 := []int{1, 87, 9}
	nums2 := []int{2, 5, 6}
	middle := findMedianSortedArrays(nums1, nums2)

	if middle != 5.5 {
		t.Error("test case failed,result=", middle)
	}

}
