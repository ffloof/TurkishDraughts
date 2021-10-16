package main

import (
	"fmt"
)

//TODO: add 

type BoardState struct {
	tiles [64]Tile
	turnTeam uint8 //0 null, 1 team white, 2 team black
}

type Tile struct {
	Team uint8 //0 empty space, 1 team white, 2 team black
	King bool
}

func (bs BoardState) GetBoardPos(x int, y int) (Tile, bool) {
	if -1 < y && y < 8 && -1 < x && y < 8 {
		return bs.tiles[(y*8)+ x], true
	}
	return Tile{ false, 0, false }, false 
}


//Get board where maximum amount of pieces were taken
func (bs BoardState) MaxTakeBoards() []BoardState {
	return []BoardState {}
}

func (bs BoardState) AllMoveBoards() []BoardState {
	for y := 0; y<8; y++ {
		for x := 0; x<8; x++ {
			tile := bs.GetBoardPos(x,y)
			if tile.Team != bs.turnTeam { continue }
			if tile.King {

			} else {

			}
		}
	}
	return []BoardState {}
}

func (bs BoardState) PlayerHasWon() uint8 { //0 = draw 1 = white wins 2 = black wins
	//If either player is out of pieces they lose
	//If a player has no playable moves they lose (checked in another part of the code)
	//If one player has a king and the other doesnt they lose
	return 0
}

func (bs BoardState) RawBoardValue() float64 {
	return 0.0
}

func (bs BoardState) BoardValue(depth int) float64 {
	//add a check for winner here
	if depth == 0 {
		//return weight of piece count in current boardstate
		return bs.RawBoardValue()
	} else {
		depth -= 1
	}

	//Loop through all possible takes -> recursively find max take
	options := bs.MaxTakeBoards()
	//If length take options == 0 move on to look through all movements 
	if len(options) == 0 {
		options = bs.AllMoveBoards()
		//Add check for no move options, which would mean other player wins
	}
	
	//Recursively find value of each boardstate
	/*
	for _, v := range options{
		BoardValue(v)
	}*/
	
	return 0.0
}



func main() {
	fmt.Println("time to find out")
}