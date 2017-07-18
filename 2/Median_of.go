package medianOf

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1, l2 := len(nums1), len(nums2)

	numbers1 := make([]float64, l1)
	numbers2 := make([]float64, l2)

	//合并两个数组
	mergeArr := make([]float64, len(nums1)+len(nums2))
	for i, val := range nums1 {
		numbers1[i] = float64(val)
	}
	for i, val := range nums2 {
		numbers2[i] = float64(val)
	}
	copy(mergeArr, numbers1)
	copy(mergeArr[len(nums1):], numbers2)

	var middle float64
	length := len(mergeArr)
	sort.Float64s(mergeArr)
	if length == 0 {
		middle = 0
	} else if length%2 == 0 {
		middle = (mergeArr[length/2] + mergeArr[length/2-1]) / 2
	} else {
		middle = float64(mergeArr[length/2])
	}
	return middle
}
