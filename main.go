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
		fmt.Println("tripping season") //TODO: either error handling or call a panic
	}
}

func (bs BoardState) Print(){
	for y:=0;y<8;y++ {
		lineStr := ""
		for x:=0;x<8;x++ {
			team := bs[(y*8)+ x].Team
			king := bs[(y*8)+ x].King
			if team == Empty {
				lineStr += "-"
			} else if team == White {
				if king {
					lineStr += "W"
				} else {
					lineStr += "w"
				}
			} else if team == Black {
				if king {
					lineStr += "B"
				} else {
					lineStr += "b"
				}
			}
		}
		fmt.Println(lineStr)
	}
}

//Get board where maximum amount of pieces were taken
func (bs BoardState) MaxTakeBoards(turnTeam Team) []BoardState {
	possibleTakeBoards := []BoardState{}
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if turnTeam != bs[(y*8)+x].Team { continue }
			if bs[(y*8)+ x].King {
				bs.FindKingTakes(x,y,0,[2]int{0,0})
			} else {
				bs.FindPawnTakes(x,y,0,[2]int{0,0})
			}
		}
	}

	return possibleTakeBoards
}

func (bs BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState) {


	return 0, []BoardState{}
}

func (bs BoardState) FindPawnTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState) {
	boards := []BoardState{ bs }
	bestTake := currentTakes
	currentTile, _ := bs.GetBoardTile(x,y) //TODO: add error checks for not on board and empty tiles

	//TODO: Add check for most takes
	//Up (white only)
	if currentTile.Team == White && !(lastDir[0] == 0 && lastDir[1] == 1) {
		newBS := bs
		jumpOverTile, onBoard1 := bs.GetBoardTile(x,y+1)
		landingTile, onBoard2 := bs.GetBoardTile(x,y+2)
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && currentTile.Team != jumpOverTile.Team {
				newBS.SetBoardTile(x,y+2, currentTile)
				newBS.SetBoardTile(x,y+1, Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})
				
				takes, possibleBoards := newBS.FindPawnTakes(x,y+2,currentTakes+1,[2]int{0,1})
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				}
				if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}
	}
	//Down (black only)
	if currentTile.Team == Black && !(lastDir[0] == 0 && lastDir[1] == -1) {
		newBS := bs
		jumpOverTile, onBoard1 := bs.GetBoardTile(x,y-1)
		landingTile, onBoard2 := bs.GetBoardTile(x,y-2)
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && currentTile.Team != jumpOverTile.Team {
				newBS.SetBoardTile(x,y-2, currentTile)
				newBS.SetBoardTile(x,y-1, Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})
				
				takes, possibleBoards := newBS.FindPawnTakes(x,y-2,currentTakes+1,[2]int{0,-1})
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				}
				if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}
	}
	//Left (either)
	if !(lastDir[0] == -1 && lastDir[1] == 0) {
		newBS := bs
		jumpOverTile, onBoard1 := bs.GetBoardTile(x-1,y)
		landingTile, onBoard2 := bs.GetBoardTile(x-2,y)
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && currentTile.Team != jumpOverTile.Team {
				newBS.SetBoardTile(x-2,y, currentTile)
				newBS.SetBoardTile(x-1,y, Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})
				
				takes, possibleBoards := newBS.FindPawnTakes(x-2,y,currentTakes+1,[2]int{-1,0})
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				}
				if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}
	} 

	//Right (either)
	if !(lastDir[0] == 1 && lastDir[1] == 0) {
		newBS := bs
		jumpOverTile, onBoard1 := bs.GetBoardTile(x+1,y)
		landingTile, onBoard2 := bs.GetBoardTile(x+2,y)
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && currentTile.Team != jumpOverTile.Team {
				newBS.SetBoardTile(x+2,y, currentTile)
				newBS.SetBoardTile(x+1,y, Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})

				takes, possibleBoards := newBS.FindPawnTakes(x+2,y,currentTakes+1,[2]int{1,0})
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				}
				if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}

	}

	return bestTake, boards
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
				moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])
				if moveTile.Team == Empty && onBoard {
					newBS := bs
					newBS.SetBoardTile(x + move[0], y + move[1], checkingTile)
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
			moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])
			if moveTile.Team == Empty && onBoard {
				newBS := bs
				if moveTile.Team == Black && y + move[1] == 0 { //Promote to king condition
					checkingTile.King = true
				} else if moveTile.Team == White && y + move[1] == 7 {
					checkingTile.King = true
				}
				newBS.SetBoardTile(x + move[0], y + move[1], checkingTile)
				newBS.SetBoardTile(x, y, Tile{})
				boards = append(boards, newBS)
			}
		}
	}
	return boards
}

func (bs BoardState) AllMoveBoards(turnTeam Team) []BoardState {
	possibleMoveBoards := []BoardState {}
	for y := 0; y<8; y++ {
		for x := 0; x<8; x++ {
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesTile(x, y, turnTeam)...)
		}
	}
	return possibleMoveBoards
}

func (bs BoardState) PlayerHasWon() Team { //0 = no winner 1 = white wins 2 = black wins
	//If either player is out of pieces they lose
	wKings := 0
	wPieces := 0

	bKings := 0
	bPieces := 0

	for _, piece := range bs {
		if piece.Team == White {
			if piece.King {
				wKings += 1
			}
			wPieces += 1
		} else if piece.Team == Black {
			if piece.King {
				bKings += 1
			}
			bPieces += 1
		}
	}

	//If a player has no moves they lose lol
	if wPieces == 0 {
		return Black
	} 

	if bPieces == 0 {
		return White
	}

	//If one player has a king and the other has one piece they lose
	if wPieces == 1 {
		if bKings > 0 {
			return Black
		}
	}

	if bPieces == 1 {
		if wKings > 0 {
			return White
		}
	}

	return Empty //No winner
	//If a player has no playable moves they lose (checked in another part of the code)
}

//Draw check would optimize returning 0 instead of worthless move searches
func (bs BoardState) PlayersDrawed() bool {
	//Check if players are in a stalemate / draw
	return false
}

func (bs BoardState) RawBoardValue() float64 { //Game is always from whites perspective
	value := 0.0
	for _, piece := range bs {
		if piece.Team == White {
			if piece.King {
				value += 5.0
			} else {
				value += 1.0
			}
		} else if piece.Team == Black {
			if piece.King {
				value -= 5.0
			} else {
				value -= 1.0
			}
		}
	}
	return value
}

//TODO: make sure to inverse board team
func (bs BoardState) BoardValue(depth int, turnTeam Team) float64 {
	//add a check for winner here
	winState := bs.PlayerHasWon()
	if winState == White {
		return 100.0
	} else if winState == Black {
		return -100.0
	}

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

//TODO: add unit tests
//TODO: try adding start move and end move table