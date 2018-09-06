package foursum

import (
	"fmt"
	"testing"
)

func Test_fourSum(t *testing.T) {

	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	fmt.Println(fourSum(nums, 0))
}
