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

func (bs BoardState) MaxTakeBoards(turnTeam Team) []BoardState {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if turnTeam != bs[(y*8)+x].Team { continue }
			var takes int
			var possibleTakeBoards []BoardState
			if bs[(y*8)+ x].King {
				takes, possibleTakeBoards = bs.FindKingTakes(x,y,0,[2]int{0,0})
			} else {
				takes, possibleTakeBoards = bs.FindPawnTakes(x,y,0)
			}
			if takes > bestTake {
					bestTake = takes
					possibleMaxTakeBoards = possibleTakeBoards
			} else if takes == bestTake {
				possibleMaxTakeBoards = append(possibleMaxTakeBoards, possibleTakeBoards...)
			}
		}
	}

	return possibleMaxTakeBoards
}

func (bs BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState) {
	boards := []BoardState{ bs }
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y) //TODO: add error checks for not on board and empty tiles

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if -lastDir[0] == direction[0] && -lastDir[1] == direction[1] { continue } //Check to not go backwards in a direction

		for i:=0;i<8;i++ {
			jumpPos := [2]int{direction[0]*i,direction[1]*i}
			jumpOverTile, onBoard := bs.GetBoardTile(jumpPos[0],jumpPos[1])
			if !onBoard { break }
			if jumpOverTile.Team == Empty { continue }
			if attackingTile.Team == jumpOverTile.Team { break }
			if attackingTile.Team != jumpOverTile.Team {
				//We have a jump
				for i=i+1;i<8;i++{
					landingPos := [2]int{direction[0]*i, direction[1]*i} 
					landingTile, onBoard := bs.GetBoardTile(landingPos[0], landingPos[1])
					if !onBoard { 
						break
					} else {
						if landingTile.Team != Empty { break }
						newBS := bs
						newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
						newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{Empty, false})
						newBS.SetBoardTile(x,y, Tile{Empty, false})

						takes, possibleBoards := newBS.FindKingTakes(landingPos[0], landingPos[1],currentTakes+1, direction)
						
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
		}		
	}

	return bestTake, boards

	return 0, []BoardState{}
}

func (bs BoardState) FindPawnTakes(x int, y int, currentTakes int) (int, []BoardState) {
	boards := []BoardState{ bs }
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y) //TODO: add error checks for not on board and empty tiles

	
	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if attackingTile.Team == White && !(move[0] == 0 && move[1] == 1) { continue } //Down (black only)
		if attackingTile.Team == Black && !(move[0] == 0 && move[1] == -1) { continue } //Up (white only)

		jumpPos := [2]int{ x+move[0],y+move[1] }
		landingPos := [2]int { x+(2*move[0]) , y+(2*move[1]) }
		jumpOverTile, onBoard1 := bs.GetBoardTile(jumpPos[0], jumpPos[1])
		landingTile, onBoard2 := bs.GetBoardTile(landingPos[0], landingPos[1])
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && attackingTile.Team != jumpOverTile.Team {
				newBS := bs
				newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
				newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})
				
				takes, possibleBoards := newBS.FindPawnTakes(landingPos[0], landingPos[1],currentTakes+1)
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

func (bs BoardState) AllMovesTile(x int, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)
	if checkingTile.King {
		for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} { //Right
			for i:=1;i<8;i++ {
				move := [2]int{direction[0]*i, direction[1]*i}
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
		for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
			if checkingTile.Team == White && !(move[0] == 0 && move[1] == 1) { continue }
			if checkingTile.Team == Black && !(move[0] == 0 && move[1] == -1) { continue }
			
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
			if turnTeam != bs[(y*8)+x].Team { continue }
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesTile(x, y)...)
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

//TODO: Draw check would optimize returning 0 instead of worthless move searches
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
func (bs BoardState) BoardValue(depth int, turnTeam Team) (float64, BoardState) {
	//add a check for winner here
	winState := bs.PlayerHasWon()
	if winState == White { 
		return 100.0, bs 
	} else if winState == Black { 
		return -100.0, bs 
	}

	if depth == 0 {
		return bs.RawBoardValue(), bs
	}

	options := bs.MaxTakeBoards(turnTeam)
	if len(options) == 0 {
		options = bs.AllMoveBoards(turnTeam)

		if len(options) == 0 { //No legal move check
			if turnTeam == White { return -100.0, bs }
			if turnTeam == Black { return 100.0, bs }
		}
	}
	

	var bestValue float64
	bestBoard := bs

	for i, branch := range options{
		if turnTeam == White {
			value, board := branch.BoardValue(depth-1, Black)
			if i==0 || value > bestValue {
				bestValue = value //White tries to maximize value
				bestBoard = board
			}
		}
		if turnTeam == Black {
			value, board := branch.BoardValue(depth-1, White)
			if i==0 || value < bestValue {
				bestValue = value //Black tries to minimize value
				bestBoard = board
			}
		}
	}
	
	return bestValue, bestBoard;
}



func main() {
	fmt.Println("time to find out")
}

//TODO: add unit tests
//TODO: try adding start move and end move table
//TODO: split into multiple files