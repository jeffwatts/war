package war

import (
	"github.com/jeffwatts/war/models"
	"testing"
)

func TestDealCards(t *testing.T) {
	game := NewGame(false)
	deck := models.NewShuffledDeck()
	game.dealCards(deck)

	expectedCardsPerPlayer := len(deck) / 2
	actualCardsA := len(game.playerA.GetHand())
	actualCardsB := len(game.playerA.GetHand())

	if actualCardsA != expectedCardsPerPlayer {
		t.Errorf("Expected playerA to have %d cards, but had %d", expectedCardsPerPlayer, actualCardsA)
	}

	if actualCardsB != expectedCardsPerPlayer {
		t.Errorf("Expected playerB to have %d cards, but had %d", expectedCardsPerPlayer, actualCardsB)
	}
}
