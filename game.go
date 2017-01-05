package war

import (
	"fmt"
	"github.com/jeffwatts/war/models"
)

type Game struct {
	playerA          models.Player
	playerB          models.Player
	commandReceiver  *CommandReceiver
	commandChannel   chan<- int
	logVerbose       bool
	playCardsChannel <-chan models.Play
}

// Creates a new 2-player game
func NewGame(logVerbose bool) *Game {
	playCardsChannel := make(chan models.Play, 8)
	playerA := models.NewPlayer("Player A", playCardsChannel)
	playerB := models.NewPlayer("Player B", playCardsChannel)
	commandChannel := make(chan int, 1)
	commandReceiver := &CommandReceiver{players: []models.Player{playerA, playerB}, commandChannel: commandChannel}

	return &Game{
		playerA:          playerA,
		playerB:          playerB,
		commandReceiver:  commandReceiver,
		commandChannel:   commandChannel,
		logVerbose:       logVerbose,
		playCardsChannel: playCardsChannel,
	}
}

func (this *Game) Play() {
	this.logv("Starting a game of war")
	deck := models.NewShuffledDeck()
	this.logv(fmt.Sprintf("Shuffled deck of cards is %v", deck))
	this.dealCards(deck)
	go this.commandReceiver.listen()
	// start the game by telling each player to play one card
	this.playNewHand()
}

// Convenience function for telling each player to play one card with no carryover
func (this *Game) playNewHand() {
	this.logv("Playing new hand")
	this.playHand(1, []models.Card{})
}

func (this *Game) playHand(numCards int, carriedOverCards []models.Card) {
	this.commandChannel <- numCards
	this.processPlay(numCards, carriedOverCards)
}

func (this *Game) dealCards(deck []models.Card) {
	for index, card := range deck {
		if index%2 == 0 {
			this.playerA.PushCards(card)
		} else {
			this.playerB.PushCards(card)
		}
	}

	this.logv(fmt.Sprintf("%s hand = %v", this.playerA.GetName(), this.playerA.GetHand()))
	this.logv(fmt.Sprintf("%s hand = %v", this.playerB.GetName(), this.playerB.GetHand()))
}

func (this *Game) processPlay(lastNumCardsRequested int, carriedOverCards []models.Card) {
	bufferA := make([]models.Card, 0, lastNumCardsRequested)
	bufferB := make([]models.Card, 0, lastNumCardsRequested)
	totalExpectedCards := 2 * lastNumCardsRequested

	for i := 0; i < totalExpectedCards; i++ {
		play := <-this.playCardsChannel
		this.logv(fmt.Sprintf("Got play = {Player: %s, Card: %v}", play.Player.GetName(), play.Card))

		// Player can't complete the play, game is over
		if play.Card == models.NilCard {
			this.endGame(play.Player)
			return
		}

		bufferA, bufferB = this.addCardToBuffer(play, bufferA, bufferB)
		this.logv(fmt.Sprintf("Buffers are now A = %v ; B = %v", bufferA, bufferB))
	}

	this.processHand(bufferA, bufferB, carriedOverCards)
}

func (this *Game) addCardToBuffer(play models.Play, bufferA, bufferB []models.Card) ([]models.Card, []models.Card) {
	if play.Player == this.playerA {
		bufferA = append(bufferA, play.Card)
	} else {
		bufferB = append(bufferB, play.Card)
	}

	return bufferA, bufferB
}

func (this *Game) processHand(playerACards, playerBCards []models.Card, carriedOverCards []models.Card) {
	lastCardA := playerACards[len(playerACards)-1]
	lastCardB := playerBCards[len(playerBCards)-1]

	// Could go all the way and push the cards through a channel
	if lastCardA > lastCardB {
		this.logv(fmt.Sprintf("%v > %v, %s wins the hand", lastCardA, lastCardB, this.playerA.GetName()))
		pushCardsToPlayer(this.playerA, playerACards, playerBCards, carriedOverCards)
		this.playNewHand()
	} else if lastCardB > lastCardA {
		this.logv(fmt.Sprintf("%v < %v, %s wins the hand", lastCardA, lastCardB, this.playerB.GetName()))
		pushCardsToPlayer(this.playerB, playerACards, playerBCards, carriedOverCards)
		this.playNewHand()
	} else { // war!
		this.logv("War!")
		carryOver := make([]models.Card, len(playerACards)+len(playerBCards)+len(carriedOverCards))
		copy(carryOver[:len(playerACards)], playerACards)
		copy(carryOver[len(playerACards):len(playerACards)+len(playerBCards)], playerBCards)
		copy(carryOver[len(playerACards)+len(playerBCards):], carriedOverCards)
		this.playHand(4, carryOver)
	}
}

func pushCardsToPlayer(player models.Player, cards ...[]models.Card) {
	for _, cardSlice := range cards {
		player.PushCards(cardSlice...)
	}
}

func (this *Game) endGame(loser models.Player) {
	var winner models.Player
	if loser == this.playerA {
		winner = this.playerB
	} else {
		winner = this.playerA
	}

	log(fmt.Sprintf("Game over! %s wins", winner.GetName()))
}

func log(msg string) {
	fmt.Println(msg)
}

func (this *Game) logv(msg string) {
	if this.logVerbose {
		log(msg)
	}
}
