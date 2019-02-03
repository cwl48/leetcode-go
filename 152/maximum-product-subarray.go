package maxproduct

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (30.75%)
 * Total Accepted:    4.1K
 * Total Submissions: 13.4K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */
func maxProduct(nums []int) int {

	result := nums[0]
	min := result
	max := result
	for i := 1; i < len(nums); i++ {

		tin := min * nums[i]
		tax := max * nums[i]

		min = getMin(getMin(tax, tin), nums[i])
		max = getMax(getMax(tax, tin), nums[i])

		result = getMax(max, result)
	}
	return result
}

func getMax(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func getMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
