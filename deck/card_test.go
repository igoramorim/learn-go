package deck

import (
	"fmt"
	"math/rand"
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

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spade as first card. Got: ", exp)
	}
}

func TestSort(t *testing.T) {
	// TODO: Test it with a custom Sort function instead of Less
	// to order the cards in a different way
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spade as first card. Got: ", exp)
	}
}

func TestShuffle(t *testing.T) {
	// Makes shuffleRand deterministic
	// Fist call to shuffleRand.Perm(52) should be?
	// [40 35 ...]
	// This overwrites the shuffleRand from the package
	shuffledRand = rand.New(rand.NewSource(0))

	original := New()
	first := original[40]
	second := original[35]

	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("Expected the first card to be %s. Got: %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s. Got: %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers. Got: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards. Got: %d", 13*4*3, len(cards))
	}
}
