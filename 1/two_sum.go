package twoSum

/*
 * [1] Two Sum
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.77%)
 * Total Accepted:    102.9K
 * Total Submissions: 240.4K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
 * 
 * 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
 * 
 * 示例:
 * 
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 * 
 * 
 */

//自己的版本 O(n^2)
// func twoSum(nums []int, target int) []int {
// 	arr := []int{}
// 	for index, val := range nums {
// 		for index1, val2 := range nums {
// 			if val+val2 == target && index != index1 {
// 				arr = []int{index, index1}
// 				return arr
// 			}
// 		}
// 	}
// 	return arr
// }

//时间复杂度小的 O(n)
func twoSum(nums []int, target int) []int {

	container := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if xy, ok := container[target-nums[i]]; ok && xy != i {
			return []int{container[target-nums[i]], i}
		}
		container[nums[i]] = i
	}
	return nil
}
