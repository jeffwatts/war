package models

import (
	"testing"
)

func TestNewShuffledDeck(t *testing.T) {
	deck := NewShuffledDeck()

	if len(deck) != 52 {
		t.Error("Expected 52-card deck, but got", deck)
	}

	for card := TWO; card <= ACE; card++ {
		foundCount := 0
		for _, val := range deck {
			if Card(card) == val {
				foundCount++
			}
		}

		if foundCount != 4 {
			t.Errorf("Expected to find 4 instances of %v but %d", card, foundCount)
		}
	}
}
