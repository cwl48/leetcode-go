package maxprofit3

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (33.42%)
 * Total Accepted:    3.2K
 * Total Submissions: 9.4K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 *
 * 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 示例 1:
 *
 * 输入: [3,3,5,0,0,3,1,4]
 * 输出: 6
 * 解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 *
 * 示例 2:
 *
 * 输入: [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
 *
 */

// O(n^2)
// func maxProfit(prices []int) int {

// 	max := 0
// 	for i := range prices {
// 		max = getMax(findMax(prices, 0, i)+findMax(prices, i, len(prices)), max)
// 	}
// 	return max
// }

// func findMax(prices []int, low int, high int) int {

// 	min := prices[low]
// 	max := 0
// 	for i := low + 1; i < high; i++ {

// 		tmp := prices[i] - min
// 		if tmp > max {
// 			max = tmp
// 		}
// 		if min > prices[i] {
// 			min = prices[i]
// 		}
// 	}
// 	return max
// }

func getMax(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// O(n)
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = getMax(buy1, -prices[i])
		sell1 = getMax(sell1, buy1+prices[i])
		buy2 = getMax(buy2, sell1-prices[i])
		sell2 = getMax(sell2, buy2+prices[i])
	}
	return getMax(0, getMax(sell1, sell2))
}
