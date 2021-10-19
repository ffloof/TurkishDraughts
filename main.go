package main

import (
	"fmt"
	"TurkishDraughts/Board"
)

func main() {
	fmt.Println("Started")
	b := board.BoardFromStr("-------- bbbbbbbb bbbbbbbb -------- -------- wwwwwwww wwwwwwww --------")
	value, nextBoard := b.BoardValue(6, board.White)
	fmt.Println(value)
	nextBoard.Print()
}

//TODO: add unit tests
//TODO: try adding start move and end move table
