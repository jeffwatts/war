package war

import (
	"testing"
)

func TestListen(t *testing.T) {
	channel := make(chan int, 1)
	underTest := CommandReceiver{commandChannel: channel}
	pushed := 5
	channel <- pushed
	underTest.listen()
}
