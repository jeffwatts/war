package models

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	playerName := "Test Player"
	playChannel := make(chan Play)
	player := NewPlayer(playerName, playChannel)

	if player.Name != playerName {
		t.Error("Expected", playerName, "got", player.Name)
	}
}

func TestPlayOneCard(t *testing.T) {
	player := new(PlayerImpl)
	playChannel := make(chan Play, 1)
	player.playChannel = playChannel
	player.PushCards(TEN, ACE)
	player.Play(1)
	played := <-playChannel
	playedCard := played.Card

	if playedCard != TEN {
		t.Error("Expected TEN to be played, but was", playedCard)
	}

	if len(player.Hand) != 1 {
		t.Error("Expected a hand with one card, got", len(player.Hand))
	}

	if player.Hand[0] != ACE {
		t.Error("Expected the hand to have an ACE, got", player.Hand)
	}
}

func TestPlayFourCards(t *testing.T) {
	player := new(PlayerImpl)
	playChannel := make(chan Play, 4)
	player.playChannel = playChannel
	player.PushCards(TEN, JACK, QUEEN, KING)
	player.Play(4)
	played := <-playChannel
	playedCard := played.Card

	if playedCard != TEN {
		t.Error("Expected TEN to be played, but was", playedCard)
	}

	if len(player.Hand) != 0 {
		t.Error("Expected a hand with no cards, got", len(player.Hand))
	}
}

func TestPlayMoreCardsThanPlayerHas(t *testing.T) {
	player := new(PlayerImpl)
	playChannel := make(chan Play, 1)
	player.playChannel = playChannel
	player.Play(1)
	played := <-playChannel

	if played.Player != player {
		t.Errorf("Expected player to be %v but was %v", player, played.Player)
	}

	if played.Card != NilCard {
		t.Error("Expected played card to be NilCard, but was", played.Card)
	}
}

func TestPushOneCard(t *testing.T) {
	player := new(PlayerImpl)
	player.PushCards(TWO)

	if len(player.Hand) != 1 {
		t.Error("Expected a hand with one card, got", len(player.Hand))
	}
}

func TestPushMultipleCards(t *testing.T) {
	player := new(PlayerImpl)
	player.PushCards(QUEEN)
	player.PushCards(TWO, ACE, SEVEN, NINE, SIX, KING)

	if len(player.Hand) != 7 {
		t.Error("Expected a hand with seven cards, got", len(player.Hand))
	}
}
