package tournament

import (
	"TurkishDraughts/Board"
)

type randomAI struct { 
	name string
}

func (rai randomAI) Play(currentBoard board.BoardState) board.BoardState {
	return board.BoardState{}
}

func (rai randomAI) GetName() string {
	return rai.name
}