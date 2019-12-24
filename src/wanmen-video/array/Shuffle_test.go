package array

import (
	"fmt"
	"testing"
)

func TestRandomShuffle_1st(t *testing.T) {
	var cards = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var ss [10][10]int
	for i := 0; i < 1000; i++ {
		r := randomShuffle_1st(cards)
		for j, c := range r {
			ss[j][c-1] += 1
		}
	}

	for _, s := range ss {
		fmt.Println(s)
	}
}
