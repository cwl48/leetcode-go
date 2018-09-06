package twoSum

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
