package board

import (
	"fmt"
	"strings"
)

//3 bit combinations for representing tiles
//0 = black, 1 = white
//0 = pawn, 1 = king
//0 = empty, 1 = piece

type TileTeam uint8
type TileKing uint8
type TileFull uint8

const (
	Empty TileFull = iota
	Filled
)

const (
	Pawn TileKing = iota
	King
)

const (
	Black TileTeam = iota
	White
)

type Tile struct {
	Team TileTeam
	King TileKing
	Full TileFull
}
type BoardState struct { 
	Team uint64 
	King uint64 
	Full uint64 
}

func (bs *BoardState) GetBoardTile(xed int, yed int) (Tile, bool) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 {
		x := uint64(xed)
		y := uint64(yed)

		team := TileTeam((bs.Team >> (y*8)+x) % 2)
		king := TileKing((bs.King >> (y*8)+x) % 2)
		full := TileFull((bs.Full >> (y*8)+x) % 2)

		return Tile{team,king,full}, true
	}
	return Tile{}, false 
}

func (bs *BoardState) SetBoardTile(xed int, yed int, t Tile) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 {
		x := uint64(xed)
		y := uint64(yed)

		team := TileTeam((bs.Team >> (y*8)+x) % 2)
		king := TileKing((bs.King >> (y*8)+x) % 2)
		full := TileFull((bs.Full >> (y*8)+x) % 2)
		
		if team != t.Team {
			bs.Team ^= (1 << (y*8)+x)
		}
		if king != t.King {
			bs.King ^= (1 << (y*8)+x)
		}
		if full != t.Full {
			bs.Full ^= (1 << (y*8)+x)
		}
	}
}

func BoardToStr(bs *BoardState) string {
	fullStr := ""
	for y:=0;y<8;y++ {
		lineStr := ""
		for x:=0;x<8;x++ {
			tile, _ := bs.GetBoardTile(x,y)
			if tile.Full == Empty {
				lineStr += "-"
			} else if tile.Team == White {
				if tile.King == King {
					lineStr += "W"
				} else {
					lineStr += "w"
				}
			} else {
				if tile.King == King {
					lineStr += "B"
				} else {
					lineStr += "b"
				}
			}
		}
		fullStr += lineStr + "\n"
	}
	fullStr = strings.TrimSpace(fullStr)
	return fullStr
}

func (bs *BoardState) Print(){
	for _, line := range strings.Fields(BoardToStr(bs)) {
		fmt.Println(line)
	} 
}

func CreateStartingBoard() BoardState {
	return BoardFromStr("-------- bbbbbbbb bbbbbbbb -------- -------- wwwwwwww wwwwwwww --------")
}

func BoardFromStr(str string) BoardState {
	rows := strings.Fields(str)
	var board BoardState

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if string(rows[y][x]) == "-" { 
				board.SetBoardTile(x,y,Tile{}) //empty
			} else if string(rows[y][x]) == "b" { 
				board.SetBoardTile(x,y,Tile{Black, Pawn, Filled}) //black pawn {TEAM KING FULL}
			} else if string(rows[y][x]) == "w" {
				board.SetBoardTile(x,y,Tile{White, Pawn, Filled}) //white pawn 
			} else if string(rows[y][x]) == "B" {
				board.SetBoardTile(x,y,Tile{Black, King, Filled}) //black king 
			} else if string(rows[y][x]) == "W" {
				board.SetBoardTile(x,y,Tile{White, King, Filled}) //white king
			} 
		}
	}
	return board
}