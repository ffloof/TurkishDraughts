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
func Search(b board.BoardState, quit chan bool, output chan PossibleMove) int {
	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	for _, branch := range options{
		go analyzeBranch(branch, board.NewTable(), output, board.Depth)
	}

	return len(options)
}

func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan PossibleMove, depth int32) {
	branch.SwapTeam()
	output <- PossibleMove {branch, branch.MinMax(depth, -board.AlphaBetaMax, board.AlphaBetaMax, table)}
	debug.FreeOSMemory()
}
