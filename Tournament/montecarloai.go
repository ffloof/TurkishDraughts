package tournament

import (
	"TurkishDraughts/Board"
)

type montecarloAI struct {
	name string
	sims int
}

func (mctsai montecarloAI) Play(currentBoard board.BoardState) board.BoardState {
	return board.MCTS(currentBoard, mctsai.sims)
}

func (mctsai montecarloAI) GetName() string {
	return mctsai.name
}