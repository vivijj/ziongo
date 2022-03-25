package core

import (
	"time"

	"github.com/vivijj/ziongo/core/state"
)

type BlockProposer struct {
	StateKeeperReq chan<- int
}

func RunProposerTask(
	miniblockInterval int,
	statekeeperReq chan<- state.ExecuteMiniBlock,
) {
	tick := time.Tick(time.Duration(miniblockInterval) * time.Second)

	for {
		select {
		case <-tick:
			statekeeperReq <- state.ExecuteMiniBlock{TimeStamp: int(time.Now().Unix())}
		}
	}
}
