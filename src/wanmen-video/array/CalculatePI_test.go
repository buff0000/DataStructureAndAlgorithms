package array

import "testing"

func TestCalculatePI(t *testing.T) {
	var arr = [...]int{-1, 1, 100, 1000, 10000, 100000, 1000000}
	for _, c := range arr {
		calculatePI(c)
	}
}
