package containerwater

/*
 * [11] Container With Most Water
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (39.75%)
 * Total Accepted:    8.8K
 * Total Submissions: 22.1K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */

/*
	写下思路
	1. 两个游标从数组两端不断往中间靠拢
	2. 计算 游标和对应的数组的乘积是否最大 即不断调整x,y  使得×*y的值最大
	3. y指代数组值，即为高度,x是两游标的差，即x轴的大小。
	4. 当所有的值都遍历之后，返回x*y最大值

*/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var (
		maxValue int               //最大的容积
		arrMax   int               //数组中对应的最大值
		left     int               //左边的游标
		right    = len(height) - 1 //右边的游标
	)

	for right > left {
		var val int
		if height[right] < height[left] {
			arrMax = height[right]
		} else {
			arrMax = height[left]
		}
		val = arrMax * (right - left)
		if val > maxValue {
			maxValue = val
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return maxValue
}
