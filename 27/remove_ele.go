package removeElement

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3}
	fmt.Println(removeElement(nums, 2))
}

/*
给定一个数组和一个值，删除该值的所有实例，并返回新的长度。

元素的顺序可以被改变，也不关心最终的数组长度。

*/

func removeElement(nums []int, val int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}
	return result
}
