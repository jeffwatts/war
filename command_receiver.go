package war

import (
	"github.com/jeffwatts/war/models"
)

type CommandReceiver struct {
	players        []models.Player
	commandChannel <-chan int
}

// Listens for commands to play cards. Should run on its own goroutine.
func (this *CommandReceiver) listen() {
	for {
		numCardsToPlay := <-this.commandChannel
		for _, player := range this.players {
			player.Play(numCardsToPlay)
		}
	}
}
