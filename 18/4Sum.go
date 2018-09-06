package foursum

import (
	"sort"
)

/*
 * [18] 4Sum
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (29.39%)
 * Total Accepted:    4.3K
 * Total Submissions: 14.5K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	arrLen := len(nums)
	result := make([][]int, 0, 0)
	temp := make([]int, 4, 4)
	var sum int
	for i := 0; i < arrLen-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < arrLen-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			low, high := j+1, arrLen-1
			for low < high {
				sum = nums[i] + nums[j] + nums[low] + nums[high]
				if sum == target {
					temp = []int{nums[i], nums[j], nums[low], nums[high]}
					result = append(result, temp)
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if sum > target {
					high--
				} else {
					low++
				}
			}
		}
	}

	return result
}
