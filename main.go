package main

import (
	"fmt"
	"TurkishDraughts/Board"
)

func main() {
	fmt.Println("Started")
	b := board.BoardFromStr("-------- bbbbbbbb bbbbbbbb -------- -------- wwwwwwww wwwwwwww --------")
	value := b.BoardValue(12, -board.AlphaBetaMax, board.AlphaBetaMax, board.White)

	/*

	options := b.MaxTakeBoards(turnTeam)
	if len(options) == 0 {
		options = b.AllMoveBoards(turnTeam)
	}

	var bestValue float64
	var bestBranch BoardState

	for i, branch := range options{
		if turnTeam == White {
			value, _ := branch.BoardValue(depth-1, Black)
			if i==0 || value >= bestValue {
				bestValue = value //White tries to maximize value
				bestBranch = branch
			}
		} else if turnTeam == Black {
			value, _ := branch.BoardValue(depth-1, White)
			if i==0 || value <= bestValue {
				bestValue = value //Black tries to minimize value
				bestBranch = branch
			}
		}
	}*/

	fmt.Println(value)
}

//TODO: add unit tests
//TODO: try adding start move and end move table
