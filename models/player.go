package models

type Player struct {
	Name                string
	Hand                []Card
	playChannel         chan<- Play
	receiveCardsChannel <-chan []Card
}

type Play struct {
	Player *Player
	Card   Card
}

func NewPlayer(name string, playChannel chan<- Play, receiveCardsChannel <-chan []Card) *Player {
	return &Player{
		Name:                name,
		Hand:                make([]Card, 0, 52),
		playChannel:         playChannel,
		receiveCardsChannel: receiveCardsChannel,
	}
}

func (this *Player) Play(numCardsToPlay int) {
	if numCardsToPlay > len(this.Hand) {
		// Play with card value = NilCard signifies that this player does not have sufficient cards and has lost
		this.playChannel <- Play{Player: this, Card: NilCard}
		return
	}

	for _, card := range this.Hand[0:numCardsToPlay] {
		this.playChannel <- Play{this, card}
	}

	this.Hand = this.Hand[numCardsToPlay:]
}

func (this *Player) PushCards(cards ...Card) {
	this.Hand = append(this.Hand, cards...)
}
