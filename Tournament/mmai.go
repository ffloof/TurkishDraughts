package tournament

import (
	"TurkishDraughts/Board"
	"runtime/debug"
	//"time"
)

type minmaxAI struct {
	name string
	table *board.TransposTable
	//4 main settings in minmax.go
	ply int32
	advanced float32
	prevIllegalBoards []board.BoardState
}

func (mmai minmaxAI) Play(currentBoard board.BoardState) board.BoardState {
	board.MaxDepth = mmai.ply - 1 //-1 Because we are searching one ply in by looping through all possibilities
	board.AdvanceWeight = mmai.advanced

	var bestEval float32
	var bestOutcome board.BoardState


	for i, consideredBoard := range mmai.filteredPossibleMoves(currentBoard)  {
		eval := consideredBoard.MinMax(0, -999.0, 999.0, mmai.table)
		if i == 0 || (currentBoard.Turn == board.White && eval > bestEval) || (currentBoard.Turn == board.Black && eval < bestEval) {
			bestEval = eval
			bestOutcome = consideredBoard
		}
	}

	return bestOutcome
}

func (mmai minmaxAI) GetName() string {
	return mmai.name
}

func (mmai minmaxAI) Update(b board.BoardState){
	mmai.prevIllegalBoards = append(mmai.prevIllegalBoards, b)
	mmai.table.Turn()
	debug.FreeOSMemory()
}

func (mmai minmaxAI) filteredPossibleMoves(currentBoard board.BoardState) []board.BoardState {
	plays := currentBoard.ValidPlays()
	for _, prevB := range mmai.prevIllegalBoards {
		for i := range plays {
			if plays[i] == prevB {
				plays = remove(plays, i)
				break
			}
		}
	}
	return plays
}

func remove(s []board.BoardState, i int) []board.BoardState {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}