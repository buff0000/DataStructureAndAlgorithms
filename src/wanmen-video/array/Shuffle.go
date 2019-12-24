package array

import (
	"math/rand"
)

//var cards = [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 'J', 'Q', 'K'}

func randomShuffle_1st(cards []byte) []byte {
	len := len(cards)
	for i, _ := range cards {
		random := i + rand.Intn(len-i)
		cards[i], cards[random] = cards[random], cards[i]
	}
	return cards
}
