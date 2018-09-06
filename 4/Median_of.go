package medianof

import (
	"sort"
)

/*
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (30.36%)
 * Total Accepted:    12.6K
 * Total Submissions: 41.5K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
 *
 * 请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
 *
 * 你可以假设 nums1 和 nums2 不同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
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
