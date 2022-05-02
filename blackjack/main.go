package main

import (
	"fmt"
	"strings"

	"deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// Ace is currently worth 1 and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		// We can't use the const values of Jack(11), Queen(12) and King(13)
		// because these 3 cards worth value 10 in Blackjack
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Gets the first Card and updates the slice 'removing' that first Card drew
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State uint8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isn't currently any player's turn!")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	// Gives capacity 5 because in Blackjack is unlikely to have more than 5
	// So we don't need to realocate the array to add more elements
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)

		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn // Player starts
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++ // TODO: Change this to the 'named' const state
	return ret
}

func EndGame(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("\n*** FINAL HANDS ***")
	fmt.Println("Player: ", ret.Player, "\nScore: ", pScore)
	fmt.Println("Dealer: ", ret.Dealer, "\nScore: ", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted!")
	case dScore > 21:
		fmt.Println("Dealer busted!")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lost!")
	case dScore == pScore:
		fmt.Println("Draw!")
	}

	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("\nPlayer: ", gs.Player)
		fmt.Println("Dealer: ", gs.Dealer.DealerString())

		fmt.Println("\nWhat will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid option: ", input)
		}
	}

	for gs.State == StateDealerTurn {
		// If dealer score <= 16, hit
		// If dealer has a soft 17, hit
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndGame(gs)
}
