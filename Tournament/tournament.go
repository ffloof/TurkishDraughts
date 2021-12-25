package tournament

import (
	"TurkishDraughts/Board"
	"fmt"
	"os"
	"time"
	"math/rand"
)

type AI interface {
	Play(board.BoardState, []board.BoardState) board.BoardState
	GetName() string
}


func Run(){
	totalTime := time.Now()
	//Round 1, mainly testing scaling performance
	bots := []AI{
		minmaxAI { "MM9", board.NewTable(6,0), 9, 0.0},
		minmaxAI { "MM8", board.NewTable(5,0), 8, 0.0},
		minmaxAI { "MM7", board.NewTable(4,0), 7, 0.0},
		minmaxAI { "MM6", board.NewTable(3,0), 6, 0.0},
		minmaxAI { "MM5", board.NewTable(2,0), 5, 0.0},
		minmaxAI { "MM4", board.NewTable(1,0), 4, 0.0},
		minmaxAI { "MM3", board.NewTable(0,0), 3, 0.0},
		montecarloAI {"MCTS2", 256},
		montecarloAI {"MCTS5", 512},
		montecarloAI {"MCTS1k", 1024},
		montecarloAI {"MCTS2k", 2048},
		montecarloAI {"MCTS4k", 4096},
		montecarloAI {"MCTS8k", 8192},
		montecarloAI {"MCTS16k", 16384},
		randomAI {"RANDOM"},
	}

	//Round 2, testing strong variants
	/*
	bots := []AI{
		minmaxAI { "MM10", board.NewTable(7,0), 10, 0.0},
		minmaxAI { "ADV9", board.NewTable(6,0), 9, 0.1},
		montecarloAI {"MCTS16k", 16384},
		montecarloAI {"MCTS64k", 65536},
		dynamicAI { "DYN11", board.NewTable(8,0), 10, 0.0 },
		dynamicAI { "COPE10", board.NewTable(7,0), 9, 0.1 },
	}

	*/

	for a := range bots {
		for b := range bots {
			roundTime := time.Now()
			OneVOne(bots[a], bots[b])
			fmt.Println("ROUND:", time.Since(roundTime).Seconds(), "TOTAL:", time.Since(totalTime).Seconds())
		}
	}
}

func OneVOne(whiteAI, blackAI AI){
	fmt.Println("Recording:", whiteAI.GetName()+"_vs_"+blackAI.GetName()+".csv")

	rand.Seed(time.Now().UnixNano())
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
			/* if b.Turn == board.White {
				fmt.Println(whiteAI.GetName(), "(AUTO WHITE)")
			} else {
				fmt.Println(blackAI.GetName(), "(AUTO BLACK)")
			} */
			b = options[0]
			f.WriteString("true,0.0," + b.ToStr() + "\n")
		} else {
			var nextBoard board.BoardState
			var duration float64

			if b.Turn == board.White {
				//Ai 1 plays
				//fmt.Println(whiteAI.GetName(), "(WHITE)")
				startTime := time.Now()
				nextBoard = whiteAI.Play(b, history)
				duration = time.Since(startTime).Seconds()
			} else {
				//Ai 2 plays
				//fmt.Println(blackAI.GetName(), "(BLACK)")
				startTime := time.Now()
				nextBoard = blackAI.Play(b, history)
				duration = time.Since(startTime).Seconds()
			}
			if nextBoard.Full == 0 { //Incase for whatever reason an invalid board gets returned by the ai, choose a random board instead
				nextBoard = options[rand.Intn(len(options)-1)]
			}
			b = nextBoard
			f.WriteString("false," + fmt.Sprintf("%.2f", duration) + "," + b.ToStr() + "\n")
		}
		//b.Print()
		//fmt.Println()
	}
}