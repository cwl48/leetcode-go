package maxsubarray

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (41.45%)
 * Total Accepted:    33.2K
 * Total Submissions: 79.9K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */
func maxSubArray(nums []int) int {

	l, max := len(nums), 0
	dp := make([]int, l)
	if l > 0 {
		dp[0] = nums[0]
		max = dp[0]
	}

	for i := 1; i < l; i++ {
		if dp[i-1] >= 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		max = getMax(max, dp[i])
	}
	return max
}

func getMax(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
