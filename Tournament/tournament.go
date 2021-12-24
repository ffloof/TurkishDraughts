package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
	//"time"
)

type AI interface {
	Play(board.BoardState) board.BoardState
	GetName() string
	Update(board.BoardState)
}


func Run(){
	OneVOne(
		minmaxAI { "MinMax2", board.NewTable(7, 0), 2, 0.0, []board.BoardState{}},
		montecarloAI { "MonteCarlo1k", 1024 })
}

func OneVOne(whiteAI, blackAI AI){
	b := board.CreateStartingBoard()
	
	//Play loop
	for {
		//Check if theres a winner and if we should stop the game loop before mcts ai violently crashes itself
		isWon, teamWon, isDraw := b.PlayerHasWon()
		if isWon || isDraw {
			if isDraw {
				//Draw
				fmt.Println("(DRAW)")
			} else if teamWon == board.White {
				//White wins
				fmt.Println(whiteAI.GetName(), "(WHITE WIN)")
			} else {
				//Black wins
				fmt.Println(blackAI.GetName(), "(BLACK WIN)")
			}
			break
		}

		//Just tells the ai, a move has happened, useful for several optimizations
		//I just use it to recycle the hash table by adjusting depth of entries
		whiteAI.Update(b)
		blackAI.Update(b)

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