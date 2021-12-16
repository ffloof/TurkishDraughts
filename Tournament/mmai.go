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
	board.MaxDepth = mmai.ply
	board.MaximumHashDepth =  mmai.maxhash
	board.TableDepthAllowedInaccuracy = mmai.inaccuracy
	board.AdvanceWeight = mmai.advanced

	_, next := currentBoard.MinMax(0, -999.0, 999.0, mmai.table)
	return next[rand.Intn(len(next))]
	return next[0]
}

func (mmai minmaxAI) GetName() string {
	return mmai.name
}

func (mmai minmaxAI) Update(){
	mmai.table.Turn()
	debug.FreeOSMemory()
}