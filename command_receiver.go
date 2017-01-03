package war

import (
	"github.com/jeffwatts/war/models"
)

type CommandReceiver struct {
	players        []*models.Player
	commandChannel <-chan int
}

func (this *CommandReceiver) listen() {
	numCardsToPlay := <-this.commandChannel
	for _, player := range this.players {
		player.Play(numCardsToPlay)
	}
}
