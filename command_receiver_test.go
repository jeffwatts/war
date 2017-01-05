package war

import (
	"testing"
	"github.com/jeffwatts/war/models"
)

type stubPlayer struct {
	playCalledTimes int
	playLastCalledWith int
}

func (s *stubPlayer) Play(numCards int) {
	s.playCalledTimes++
	s.playLastCalledWith = numCards
}

func (s stubPlayer) PushCards(cs ...models.Card) {}

func (s stubPlayer) GetName() string {
	return "stub"
}

func (s stubPlayer) GetHand() []models.Card {
	return []models.Card{}
}


func TestListen(t *testing.T) {
	channel := make(chan int)
	player := new(stubPlayer)
	underTest := CommandReceiver{players: []models.Player{player}, commandChannel: channel}
	pushed := 4
	go underTest.listen()
	channel <- pushed

	if player.playCalledTimes != 1 {
		t.Errorf("Expected Play() to be called 1 time, but was %d", player.playCalledTimes)
	}

	if player.playLastCalledWith != pushed {
		t.Errorf("Expected Play() to be called with %d but was %d", pushed, player.playLastCalledWith)
	}
}
