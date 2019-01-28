package search81

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (31.86%)
 * Total Accepted:    3.2K
 * Total Submissions: 10.1K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 */
func search(nums []int, target int) bool {

	low, high, mid := 0, len(nums)-1, 0

	for low <= high {

		mid = low + (high-low)/2

		if nums[low] == target || nums[mid] == target || nums[high] == target {
			return true
		}

		if nums[low] < nums[mid] {
			if target > nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[low] > nums[mid] {
			if target < nums[low] && target > nums[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			low++
		}
	}
	return false
}
