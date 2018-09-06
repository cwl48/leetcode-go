package sumoftwointegers

/*
 * [371] Sum of Two Integers
 *
 * https://leetcode-cn.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (54.69%)
 * Total Accepted:    3.2K
 * Total Submissions: 5.8K
 * Testcase Example:  '1\n2'
 *
 * 不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
 * 
 * 示例 1:
 * 
 * 输入: a = 1, b = 2
 * 输出: 3
 * 
 * 
 * 示例 2:
 * 
 * 输入: a = -2, b = 3
 * 输出: 1
 * 
 */

func getSum(a int, b int) int {

	var _a int
	var _b int
	for b != 0 {
		_a = a ^ b
		_b = (a & b) << 1
		a = _a
		b = _b
	}
	return a

}
