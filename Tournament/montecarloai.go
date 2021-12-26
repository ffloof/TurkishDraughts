package tournament

import (
	"TurkishDraughts/Board"
)

type montecarloAI struct {
	name string
	sims int
}

func (mctsai montecarloAI) Play(currentBoard board.BoardState, prevIllegalBoards []board.BoardState) board.BoardState {
	board.MonteIllegalBoards = prevIllegalBoards
	return board.MCTS(currentBoard, mctsai.sims)
}

func (mctsai montecarloAI) GetName() string {
	return mctsai.name
}