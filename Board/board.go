package board

import (
	"fmt"
	"strings"
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

func (bs *BoardState) GetBoardTile(x int, y int) (Tile, bool) {
	if -1 < x && x < 8 && -1 < y && y < 8 {
		return bs[(y*8)+ x], true
	}
	return Tile{}, false 
}

func (bs *BoardState) SetBoardTile(x int, y int, t Tile) {
	if -1 < x && x < 8 && -1 < y && y < 8 {
		bs[(y*8)+ x] = t
	} else {
		fmt.Println("tripping season") //TODO: either error handling or call a panic
	}
}

func (bs *BoardState) Print(){
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

func CreateStartingBoard() BoardState {
	var bs BoardState
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if y == 1 || y == 2 { bs[(y*8)+ x].Team = Black }
			if y == 5 || y == 6 { bs[(y*8)+ x].Team = White }
		}		
	}
	return bs
}

func BoardFromStr(str string) BoardState {
	rows := strings.Fields(str)
	var board BoardState

	//TODO: error out of bound checks
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if string(rows[y][x]) == "-" { 
				board.SetBoardTile(x,y,Tile{}) //empty
			} else if string(rows[y][x]) == "b" { 
				board.SetBoardTile(x,y,Tile{Black, false}) //black pawn
			} else if string(rows[y][x]) == "w" {
				board.SetBoardTile(x,y,Tile{White, false}) //white pawn 
			} else if string(rows[y][x]) == "B" {
				board.SetBoardTile(x,y,Tile{Black, true}) //black king 
			} else if string(rows[y][x]) == "W" {
				board.SetBoardTile(x,y,Tile{White, true}) //white king
			} 
		}
	}
	return board
}