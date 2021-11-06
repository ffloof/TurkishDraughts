package main

import (
	"TurkishDraughts/Network"
	"TurkishDraughts/Board"
)

func main() {
	//network.Init()
	b := board.CreateStartingBoard()

	//Ai plays against itself for my amusement
	for true {
		b = *(network.Analyze(b, 10))
		b.Print()
		if b.RawBoardValue() >= board.WinWeight { break }
		if b == board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- --------") { break }
	}
}