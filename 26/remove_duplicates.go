package removeDuplicates

import (
	"fmt"
)

func main() {

	nums := []int{1, 1, 5, 6, 6, 6}
	fmt.Println(removeDuplicates(nums))

}

/*
 翻译 给定一个已排序的数组，删除重复的元素，这样每个元素只出现一次，并且返回新的数组长度。
 不允许为另一个数组使用额外的空间，你必须就地以常量空间执行这个操作。
 例如， 给定输入数组为 [1,1,2] 你的函数应该返回length = 2， 其前两个元素分别是1和2。
*/
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}
