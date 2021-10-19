package main

import (
	"fmt"
	"TurkishDraughts/Board"
)

func main() {
	x := board.CreateStartingBoard()
	x.Print()
}

//TODO: add unit tests
//TODO: try adding start move and end move table
