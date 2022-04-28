package deck

import (
	"fmt"
	"testing"
)

// Examples tests runs with 'go test'
func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Five, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Five of Clubs
	// Queen of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}
