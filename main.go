package main

import (
	"fmt"
	"TurkishDraughts/Board"
)

func main() {
	fmt.Println("Started")
	b := board.BoardFromStr("-------- --W---w- -------- -------- -------- -------- -------- --------")
	
	options := b.AllMoveBoards(board.White)

	for i, branch := range options{
		fmt.Println()
		fmt.Println(i)
		branch.Print()
	}
}

//TODO: add unit tests
//TODO: try adding start move and end move table