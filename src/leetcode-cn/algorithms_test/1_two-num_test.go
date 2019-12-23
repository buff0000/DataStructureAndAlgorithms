package algorithms

import (
	"testing"
	"fmt"
	"leetcode-cn/algorithms"
)

func TestTwoSum(t *testing.T) {
	numsarr := [][]int{
		[]int{2, 5, 7, 19, 55},
		[]int{2, 5},
		[]int{2, 7},
		[]int{2, 9},
		[]int{2},
		[]int{},
	}

	for _, nums := range numsarr {
		res := algorithms.TwoSum(nums, 9)
		fmt.Println(res)
	}
}
