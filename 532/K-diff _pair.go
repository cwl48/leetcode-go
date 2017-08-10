package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 5, 4}
	fmt.Println(findPairs(nums, 2))
}

func findPairs(nums []int, k int) int {

	if k < 0 || len(nums) < 0 || nums == nil {
		return 0
	}
	//构建一个map，存放(数字,个数) 值
	contain := make(map[int]int)
	count := 0 //对数量
	for _, val := range nums {
		contain[val]++
	}
	for entry, val := range contain {
		if k == 0 {
			if val >= 2 {
				count++
			}
		} else {
			if _, ok := contain[entry+k]; ok {

				count++
			}
		}
	}
	return count
}
