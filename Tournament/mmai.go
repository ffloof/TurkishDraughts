package tournament

import (
	"TurkishDraughts/Board"
	"runtime/debug"
	"math/rand"
	"time"
)

type minmaxAI struct {
	name string
	table *board.TransposTable
	//4 main settings in minmax.go
	ply int32
	advanced float32
	maxhash int32
	inaccuracy int32

}

func (mmai minmaxAI) Play(currentBoard board.BoardState) board.BoardState {
	rand.Seed(time.Now().UnixNano())
	board.MaxDepth = mmai.ply - 1 //-1 Because we are searching one ply in by looping through all possibilities
	board.MaximumHashDepth =  mmai.maxhash
	board.TableDepthAllowedInaccuracy = mmai.inaccuracy
	board.AdvanceWeight = mmai.advanced

	var bestEval float32
	var bestOutcome board.BoardState

	for i, consideredBoard := range currentBoard.ValidPlays() {
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

func (mmai minmaxAI) Update(){
	mmai.table.Turn()
	debug.FreeOSMemory()
}