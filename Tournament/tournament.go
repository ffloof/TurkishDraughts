package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
	"os"
	"time"
)

type AI interface {
	Play(board.BoardState, []board.BoardState) board.BoardState
	GetName() string
}


func Run(){
	OneVOne(
		minmaxAI { "MinMax9", board.NewTable(7, 0), 9, 0.0},
		minmaxAI { "MinMax10", board.NewTable(7, 0), 10, 0.0})
}

func OneVOne(whiteAI, blackAI AI){
	//Setup logging
	f, err := os.Create(whiteAI.GetName()+"_vs_"+blackAI.GetName()+".csv")
	if err != nil { fmt.Println(err) }
	defer f.Close()

	//Header bool, float64, string
	f.WriteString("AUTO,TIME,BOARD\n")

	//Setup board
	history := []board.BoardState{}
	b := board.CreateStartingBoard()

	//Play loop
	for {
		history = append(history, b)

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

		options := b.ValidPlays()
		if len(options) == 1 {
			if b.Turn == board.White {
				fmt.Println(whiteAI.GetName(), "(AUTO WHITE)")
			} else {
				fmt.Println(blackAI.GetName(), "(AUTO BLACK)")
			}
			b = options[0]
			f.WriteString("true,0.0," + b.ToStr() + "\n")
		} else {
			startTime := time.Now()
			if b.Turn == board.White {
				//Ai 1 plays
				fmt.Println(whiteAI.GetName(), "(WHITE)")
				b = whiteAI.Play(b, history)
			} else {
				//Ai 2 plays
				fmt.Println(blackAI.GetName(), "(BLACK)")
				b = blackAI.Play(b, history)
			}
			duration := time.Since(startTime).Seconds()
			f.WriteString("false," + fmt.Sprintf("%.2f", duration) + "," + b.ToStr() + "\n")
		}
		b.Print()
		fmt.Println()
	}
}