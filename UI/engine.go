package ui

import (
	"runtime/debug"
	"TurkishDraughts/Board"
)

type PossibleMove struct {
	board board.BoardState
	value float32 
}

//Two channels one for results back, and one for if it should quit searching
func Search(b board.BoardState, output chan PossibleMove) int {
	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	for _, branch := range options{
		branch.SwapTeam()
		go analyzeBranch(branch, board.NewTable(), output)
	}

	return len(options)
}

func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan PossibleMove) {
	output <- PossibleMove {branch, branch.MinMax(0, -board.AlphaBetaMax, board.AlphaBetaMax, table)}
	debug.FreeOSMemory()
}
