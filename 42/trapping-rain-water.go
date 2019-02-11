package trap42

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (40.13%)
 * Total Accepted:    7.7K
 * Total Submissions: 19.1K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */
func trap(height []int) int {

	sum := 0
	low, high := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for low <= high {

		leftMax = getMax(height[low], leftMax)
		rightMax = getMax(height[high], rightMax)

		if leftMax <= rightMax {
			sum += (leftMax - height[low])
			low++
		} else {
			sum += (rightMax - height[high])
			high--
		}

	}
	return sum
}

func getMax(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
