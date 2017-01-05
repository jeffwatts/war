package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Card uint8

// No suits because they don't matter in war
const (
	// NilCard represents an empty hand
	NilCard = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

func (c Card) String() string {
	switch c {
	case NilCard:
		return "NIL"
	case JACK:
		return "Jack"
	case QUEEN:
		return "Queen"
	case KING:
		return "King"
	case ACE:
		return "Ace"
	default:
		return fmt.Sprintf("%d", c)
	}
}

// Creates a standard, shuffled 52-card deck
func NewShuffledDeck() []Card {
	deck := make([]Card, 52)
	rand.Seed(time.Now().UTC().UnixNano())
	ordering := rand.Perm(52)
	var curCard Card
	curCard = TWO
	for _, targetIndex := range ordering {
		deck[targetIndex] = curCard
		curCard = (curCard + 1) % 14
		if curCard == NilCard {
			curCard = TWO
		}
	}

	return deck
}
