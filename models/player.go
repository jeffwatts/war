package models

type Player interface {
	Play(numCards int)
	PushCards(cards ...Card)
	GetName() string
	GetHand() []Card
}

type PlayerImpl struct {
	Name                string
	Hand                []Card
	playChannel         chan<- Play
}

type Play struct {
	Player Player
	Card   Card
}

func NewPlayer(name string, playChannel chan<- Play) *PlayerImpl {
	return &PlayerImpl{
		Name:                name,
		Hand:                make([]Card, 0, 52),
		playChannel:         playChannel,
	}
}

func (this *PlayerImpl) Play(numCardsToPlay int) {
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

func (this *PlayerImpl) PushCards(cards ...Card) {
	this.Hand = append(this.Hand, cards...)
}

func (this *PlayerImpl) GetName() string {
	return this.Name
}

func (this *PlayerImpl) GetHand() []Card {
	return this.Hand
}
