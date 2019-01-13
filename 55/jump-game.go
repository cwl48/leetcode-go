package canjump

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (31.43%)
 * Total Accepted:    9K
 * Total Submissions: 28.4K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 *
 */
func canJump(nums []int) bool {

	r, l, maxR := 0, 0, 0

	for r < (len(nums) - 1) {
		for i := l; i <= r; i++ {
			maxR = max(maxR, nums[i]+i)
		}
		l++
		r = maxR
		if l > (len(nums) - 1) {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// the faster one
// func canJump(nums []int) bool {
// 	last := 0
    
//     for i, num := range nums {
//         if i > last {
//             return false
//         }
        
//         if i + num > last {
//             last = i + num
//         }
        
//         if last >= len(nums) - 1{
//             return true
//         }
//     }
//     return true
// }
