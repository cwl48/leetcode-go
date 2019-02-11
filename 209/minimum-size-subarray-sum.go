package minsubarraylen

import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (35.87%)
 * Total Accepted:    5.6K
 * Total Submissions: 15.6K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例:
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 */
func minSubArrayLen(s int, nums []int) int {

	start, end, curSum, min, l := 0, 0, 0, math.MaxInt64, len(nums)

	for start < l {

		if curSum < s && end < l {
			curSum += nums[end]
			end++
		} else if curSum >= s {
			min = getMin(min, end-start)
			curSum -= nums[start]
			start++
		} else {
			break
		}
	}
	if min == math.MaxInt64 {
		return 0
	} else {
		return min
	}
}
func getMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
