package array

import (
	"fmt"
	"math"
)

/**
利用公式pi=4-4/3+4/5-4/7+…计算pi值，输入参数为计算次数
*/
func calculatePI(c int) {
	if c <= 0 {
		fmt.Println("计算次数不能小于等于零")
	}
	var pi float64
	for i := 1; i <= c; i++ {
		pi += math.Pow(-1, float64(i+1)) * 4 / float64(2*i-1)
	}
	fmt.Println("计算结果为：", pi)
}
