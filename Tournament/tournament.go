package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
	"time"
	"math/rand"
	"runtime/debug"
)

type AI interface {
	Play(board.BoardState) board.BoardState
	GetName() string
	Update()
}

type minmaxAI struct {
	name string
	table *board.TransposTable
	//4 main settings in minmax.go
	ply int32
	advanced float32
	maxhash int32
	inaccuracy int32

}

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

func (mctsai montecarloAI) Update(){}

func (mmai minmaxAI) Play(currentBoard board.BoardState) board.BoardState {
	board.MaxDepth = mmai.ply
	board.MaximumHashDepth =  mmai.maxhash
	board.TableDepthAllowedInaccuracy = mmai.inaccuracy
	board.AdvanceWeight = mmai.advanced

	_, next := currentBoard.MinMax(0, -999.0, 999.0, mmai.table) //TODO: implement table recycling
	return next[rand.Intn(len(next))]
}

func (mmai minmaxAI) GetName() string {
	return mmai.name
}

func (mmai minmaxAI) Update(){
	mmai.table.Turn()
	debug.FreeOSMemory()
}

func Run(){
	OneVOne(
		minmaxAI { "MinMax10", board.NewTable(), 10, 0.0, 8, 0},
		minmaxAI { "MinMax9", board.NewTable(), 9, 0.0, 7, 0})
}

func OneVOne(whiteAI, blackAI AI){
	rand.Seed(time.Now().UnixNano())

	b := board.CreateStartingBoard()
	for {
		//Just tells the ai, a move has happened, useful for several optimizations
		//I just use it to recycle the hash table by adjusting depth of entries
		whiteAI.Update()
		blackAI.Update()

		options := b.ValidPlays()
		if len(options) == 1 {
			if b.Turn == board.White {
				fmt.Println(whiteAI.GetName(), "(AUTO WHITE)")
			} else {
				fmt.Println(blackAI.GetName(), "(AUTO BLACK)")
			}
			b = options[0]
		} else {
			if b.Turn == board.White {
				//Ai 1 plays
				b = whiteAI.Play(b)
				fmt.Println(whiteAI.GetName(), "(WHITE)")
			} else {
				//Ai 2 plays
				b = blackAI.Play(b)
				fmt.Println(blackAI.GetName(), "(BLACK)")
			}
		}
		b.Print()
		fmt.Println()
	}
}