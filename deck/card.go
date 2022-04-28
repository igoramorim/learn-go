//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spade   Suit = iota // iota gives Spade value 0 and the following Suits are incremented by 1
	Diamond             // 1
	Club                // 2
	Heart               // 3
	Joker               // 4 It is a special case, not really a suit
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota // Starting with blank identifier so each Rank gets the corresponding value (Ace = 1, Two = 2 ...)
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}
