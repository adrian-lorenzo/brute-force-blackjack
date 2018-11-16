package blackjack

import (
	"fmt"
	"time"

	"github.com/adrianlorenzodev/brute-force/go/iterator"
)

// ProbabilityCheck - calculates the probability of not passing in blackjack
// cards []int - Deck of cards, without the card the player has
// cardSelected int - The card that the player has
// numCards int - Number of cards that the player wants
func ProbabilityCheck(cards []int, cardSelected, numCards int) {
	iter := iterator.InitIterator(len(cards), numCards)

	tries := 0.00
	wins := 0.00

	start := time.Now().UnixNano() / int64(time.Millisecond)

	for iter.HasNext() {
		comb := iter.Next()
		result := 0
		for _, number := range comb {
			result += cards[number]
		}

		if (result + cardSelected) <= 21 {
			wins++
		}

		tries++
	}

	fmt.Printf("The probability of not passing in blackjack with a %d choosing %d cards is of %f%%. TIME: %dms\n",
		cardSelected, numCards, ((wins / tries) * 100), (time.Now().UnixNano()/int64(time.Millisecond))-start)
}
