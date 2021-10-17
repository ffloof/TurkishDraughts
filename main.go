package main

import (
	"fmt"
)

type Team uint8 

const (
	Empty Team = iota
	White
	Black
)

type Tile struct {
	Team Team //0 empty space, 1 team white, 2 team black
	King bool
}

type BoardState [64]Tile

func (bs BoardState) GetBoardTile(x int, y int) (Tile, bool) {
	if -1 < y && y < 8 && -1 < x && y < 8 {
		return bs[(y*8)+ x], true
	}
	return Tile{}, false 
}

func (bs BoardState) SetBoardTile(x int, y int, t Tile) {
	if -1 < y && y < 8 && -1 < x && y < 8 {
		bs[(y*8)+ x] = t
	} else {
		fmt.Println("tripping season")
	}
}


//Get board where maximum amount of pieces were taken
func (bs BoardState) MaxTakeBoards(turnTeam Team) []BoardState {
	return []BoardState {}
}

func (bs BoardState) AllMovesTile(x int, y int, turnTeam Team) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)
	if checkingTile.Team != turnTeam { return boards }
	if checkingTile.King {
		var moves = [4][7][2]int{
			{	
				{0,-1}, //Up
				{0,-2},
				{0,-3},
				{0,-4},
				{0,-5},
				{0,-6},
				{0,-7},
			}, {
				{0,1}, //Down
				{0,2},
				{0,3},
				{0,4},
				{0,5},
				{0,6},
				{0,7},
			}, {
				{-1,0}, //Left
				{-2,0},
				{-3,0},
				{-4,0},
				{-5,0},
				{-6,0},
				{-7,0},
			}, {
				{1,0},  //Right
				{2,0},
				{3,0},
				{4,0},
				{5,0},
				{6,0},
				{7,0},
			},
		}
		for _, direction := range moves {
			for _, move := range direction {
				moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[0])
				if moveTile.Team == Empty && onBoard {
					newBS := bs
					newBS.SetBoardTile(x + move[0], y+ move[0], checkingTile)
					newBS.SetBoardTile(x, y, Tile{})
					boards = append(boards, newBS)
				} else {
					break //Stops going in this direction after it hits wall or piece
				}
			}
		}
	} else {
		var moves [3][2]int
		if checkingTile.Team == White {
			moves = [3][2]int{ //TODO: make these constants
				{0,-1}, //Up
				{-1,0}, //Left
				{1,0},  //Right
			}
		} else if checkingTile.Team == Black {
			moves = [3][2]int{
				{0,1}, //Down
				{-1,0}, //Left
				{1,0},  //Right
			}
		}

		for _, move := range moves {
			moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[0])
			if moveTile.Team == Empty && onBoard {
				newBS := bs
				newBS.SetBoardTile(x + move[0], y+ move[0], checkingTile)
				newBS.SetBoardTile(x, y, Tile{})
				boards = append(boards, newBS)
			}
		}
	}
	return boards
}

func (bs BoardState) AllMoveBoards(turnTeam Team) []BoardState {
	allBoards := []BoardState {}
	for y := 0; y<8; y++ {
		for x := 0; x<8; x++ {
			allBoards = append(allBoards, bs.AllMovesTile(x, y, turnTeam)...)
		}
	}
	return allBoards
}

func (bs BoardState) PlayerHasWon() Team { //0 = draw 1 = white wins 2 = black wins
	//If either player is out of pieces they lose
	//If a player has no playable moves they lose (checked in another part of the code)
	//If one player has a king and the other doesnt they lose
	return Empty
}

func (bs BoardState) PlayersDrawed() bool {
	//Check if players are in a stalemate / draw
	return false
}

func (bs BoardState) RawBoardValue() float64 {
	return 0.0
}

//TODO: make sure to inverse board team
func (bs BoardState) BoardValue(depth int, turnTeam Team) float64 {
	//add a check for winner here
	if depth == 0 {
		//return weight of piece count in current boardstate
		return bs.RawBoardValue()
	} else {
		depth -= 1
	}

	//Loop through all possible takes -> recursively find max take
	options := bs.MaxTakeBoards(turnTeam)
	//If length take options == 0 move on to look through all movements 
	if len(options) == 0 {
		options = bs.AllMoveBoards(turnTeam)
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