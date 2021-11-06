package main

import (
	//"TurkishDraughts/Network"
	"TurkishDraughts/Board"
	"fmt"
)



func main() {
	//network.Init()

	b := board.BoardFromStr("-------- -------- -b------ --b-b-W- b------- -------- -------- --------")
	b.Turn = board.White
	for _, b2 := range b.MaxTakeBoards() {
		fmt.Println(" ")
		b2.Print()
	}
	/*
	//Ai plays against itself for my amusement
	b := board.CreateStartingBoard()
	for true {
		b = *(network.Analyze(b, 10))
		b.Print()
		if b.RawBoardValue() >= board.WinWeight { break }
		if b == board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- --------") { break }
	}*/
}