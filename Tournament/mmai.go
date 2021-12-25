package tournament

import (
	"TurkishDraughts/Board"
	"runtime/debug"
	//"time"
)

type minmaxAI struct {
	name string
	table *board.TransposTable
	ply int32
	advanced float32
}

func (mmai minmaxAI) Play(currentBoard board.BoardState, prevIllegalBoards []board.BoardState) board.BoardState {
	mmai.table.Turn()

	board.MaxDepth = mmai.ply - 1 //-1 Because we are searching one ply in by looping through all possibilities
	board.AdvanceWeight = mmai.advanced

	var bestEval float32
	var bestOutcome board.BoardState

	plays := currentBoard.ValidPlays()
	for _, prevB := range prevIllegalBoards {
		for i := range plays {
			if plays[i] == prevB {
				plays = remove(plays, i)
				break
			}
		}
	}


	for i, consideredBoard := range plays {
		eval := consideredBoard.MinMax(0, -999.0, 999.0, mmai.table)
		if i == 0 || (currentBoard.Turn == board.White && eval >= bestEval) || (currentBoard.Turn == board.Black && eval <= bestEval) {
			if bestEval == eval {
				tempAW := -board.AdvanceWeight + 0.1
				//Tie breaker function
				if (currentBoard.Turn == board.White && consideredBoard.RawBoardValue(tempAW) >= bestOutcome.RawBoardValue(tempAW)) || (currentBoard.Turn == board.Black && consideredBoard.RawBoardValue(tempAW) <= bestOutcome.RawBoardValue(tempAW)){
					bestOutcome = consideredBoard
				}
			} else {
				bestEval = eval
				bestOutcome = consideredBoard
			}
		}
	}

	mmai.table.Turn()
	debug.FreeOSMemory()
	return bestOutcome
}

func (mmai minmaxAI) GetName() string {
	return mmai.name
}


func remove(s []board.BoardState, i int) []board.BoardState {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}