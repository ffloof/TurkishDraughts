package board

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