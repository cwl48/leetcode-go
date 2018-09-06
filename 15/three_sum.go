package threeSum

import (
	"sort"
)

/*
 * [15] 3Sum
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (15.53%)
 * Total Accepted:    11.9K
 * Total Submissions: 76.6K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

/*
	写下思路
	要从一个数组中获取三个数，使得x+y+z=0
	先对数组排序，for循环遍历数组，用0-nums【i】得到sum
	再对 数组i下标之后的数进行首尾遍历，当[low]+[high]==sum,则i,low,high就是一个答案
	当low和high是一个答案:对于low之后和【low】值相等的数进行跳过或者high之前和[high]值相等的数进行跳过
	如果不是答案，当low+high>sum,则low--,否则high++
*/
func threeSum(nums []int) [][]int {

	array := make([][]int, 0, 0)
	arr := make([]int, 0, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		low, high := i+1, len(nums)-1
		sum := 0 - nums[i]
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			for low < high {
				if sum == nums[low]+nums[high] {
					arr = []int{nums[i], nums[low], nums[high]}
					array = append(array, arr)
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if sum > nums[low]+nums[high] {
					low++
				} else {
					high--
				}
			}
		}

	}
	return array
}
