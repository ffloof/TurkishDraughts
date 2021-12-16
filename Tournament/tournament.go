package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
	//"time"
)

type AI interface {
	Play(board.BoardState) board.BoardState
	GetName() string
	Update()
}


func Run(){
	OneVOne(
		minmaxAI { "MinMax10", board.NewTable(), 10, 0.0, 0, 0},
		minmaxAI { "MinMax9", board.NewTable(), 9, 0.0, 0, 0})
}

func OneVOne(whiteAI, blackAI AI){
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
		//b.Print() //TODO: uncomment this
		fmt.Println()
	}
}