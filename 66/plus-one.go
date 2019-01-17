package plusone

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (36.29%)
 * Total Accepted:    31.5K
 * Total Submissions: 86.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */
func plusOne(digits []int) []int {

	times := 0
	m := len(digits)
	digits[m-1] = digits[m-1] + 1
	if digits[m-1] == 10 {
		times++
		digits[m-1] = 0
		for i := m - 2; i >= 0; i-- {
			digits[i] = digits[i] + 1
			if digits[i] < 10 {
				break
			} else {
				digits[i] = 0
				times++
			}
		}

		// 是否超出位数
		if times == m {
			// 重新生成digits
			arr := make([]int, m+1)
			arr[0] = 1
			for j := 0; j < m; j++ {
				arr[j+1] = digits[j]
			}
			digits = arr
		}
	}
	return digits
}
