package tournament

import (
	"TurkishDraughts/Board"
	"runtime/debug"
)

type dynamicAI struct {
	name string
	table *board.TransposTable
	ply int32
	advanced float32
}

func (dmmai dynamicAI) Play(currentBoard board.BoardState) board.BoardState {
	dmmai.table.Turn()
	board.AdvanceWeight = dmmai.advanced

	var bestEval float32
	var bestOutcome board.BoardState

	plays := currentBoard.ValidPlays()
	for _, prevB := range board.IllegalBoards {
		for i := range plays {
			if plays[i] == prevB {
				plays = remove(plays, i)
				break
			}
		}
	}

	board.MaxDepth = dmmai.ply - (int32(len(plays))/10) //Adjust depth based off amount of plays that will be searched

	for i, consideredBoard := range plays {
		eval := consideredBoard.MinMax(0, -999.0, 999.0, dmmai.table)
		if i == 0 || (currentBoard.Turn == board.White && eval >= bestEval) || (currentBoard.Turn == board.Black && eval <= bestEval) {
			if bestEval == eval {
				//Tie breaker functionality
				tempAW := 0.1 - board.AdvanceWeight
				if (currentBoard.Turn == board.White && consideredBoard.RawBoardValue(tempAW) >= bestOutcome.RawBoardValue(tempAW)) || (currentBoard.Turn == board.Black && consideredBoard.RawBoardValue(tempAW) <= bestOutcome.RawBoardValue(tempAW)){
					bestOutcome = consideredBoard
				}
			} else {
				bestEval = eval
				bestOutcome = consideredBoard
			}
		}
	}

	dmmai.table.Turn()
	debug.FreeOSMemory()
	return bestOutcome
}

func (dmmai dynamicAI) GetName() string {
	return dmmai.name
}