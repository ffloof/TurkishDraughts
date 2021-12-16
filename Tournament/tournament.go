package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
)

type AI interface {
	Play(board.BoardState) board.BoardState
	GetName() string
}

type minmaxAI struct {
	name string
	table board.TransposTable
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

func (mmai minmaxAI) Play(currentBoard board.BoardState) board.BoardState {
	board.MaxDepth = mmai.ply
	board.MaximumHashDepth =  mmai.maxhash
	board.TableDepthAllowedInaccuracy = mmai.inaccuracy
	board.AdvanceWeight = mmai.advanced

	_, next := currentBoard.MinMax(0, -999.0, 999.0, board.NewTable()) //TODO: implement table recycling
	return *next
}

func (mmai minmaxAI) GetName() string {
	return mmai.name
}

func Run(){
	var AI1 AI = minmaxAI { "MinMax10", *(board.NewTable()), 10, 0.0, 0, 0}
	var AI2 AI = minmaxAI { "MinMax10", *(board.NewTable()), 10, 0.0, 0, 0}

	b := board.CreateStartingBoard()
	for {
		options := b.ValidPlays()
		if len(options) == 1 {
			if b.Turn == board.White {
				fmt.Println(AI1.GetName(), "(AUTO WHITE)")
			} else {
				fmt.Println(AI2.GetName(), "(AUTO BLACK)")
			}
			b = options[0]
		} else {
			if b.Turn == board.White {
				//Ai 1 plays
				b = AI1.Play(b)
				fmt.Println(AI1.GetName(), "(WHITE)")
			} else {
				//Ai 2 plays
				b = AI2.Play(b)
				fmt.Println(AI2.GetName(), "(BLACK)")
			}
		}
		b.Print()
		fmt.Println()
	}
}