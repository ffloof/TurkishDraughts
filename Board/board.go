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
	Turn TileTeam
	Team uint64 
	King uint64 
	Full uint64 
}

func (bs *BoardState) GetBoardTile(xed int, yed int) (Tile, bool) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 {
		x := uint64(xed)
		y := uint64(yed)

		team := (bs.Team >> (y*8+x)) & 1
		king := (bs.King >> (y*8+x)) & 1
		full := (bs.Full >> (y*8+x)) & 1

		return Tile{TileTeam(team),TileKing(king),TileFull(full)}, true
	}
	return Tile{}, false 
}

func (bs *BoardState) SetBoardTile(xed int, yed int, t Tile) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 {
		x := uint64(xed)
		y := uint64(yed)

		team := TileTeam((bs.Team >> (y*8+x)) & 1)
		king := TileKing((bs.King >> (y*8+x)) & 1)
		full := TileFull((bs.Full >> (y*8+x)) & 1)
		
		if team != t.Team {
			bs.Team ^= 1 << (y*8+x)
		}
		if king != t.King {
			bs.King ^= 1 << (y*8+x)
		}
		if full != t.Full {
			bs.Full ^= 1 << (y*8+x)
		}
	}
}

func (bs *BoardState) SwapTeam(){
	bs.TryPromotion()
	if bs.Turn == White {
		bs.Turn = Black
	} else {
		bs.Turn = White
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
	b := BoardFromStr("-------- bbbbbbbb bbbbbbbb -------- -------- wwwwwwww wwwwwwww --------")
	b.Turn = White
	return b
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

func (bs *BoardState) PlayerHasWon() (bool, TileTeam) { 
	//If either player is out of pieces they lose
	var wKings uint8 = 0
	var wPieces uint8 = 0

	var bKings uint8 = 0
	var bPieces uint8 = 0

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty { continue }
		if piece.Team == White {
			wKings += uint8(piece.King)
			wPieces += 1
		} else {
			bKings += uint8(piece.King)
			bPieces += 1
		}		
	}

	//If a player has no moves they lose lol
	if wPieces == 0 {
		return true, Black
	} 

	if bPieces == 0 {
		return true, White
	}

	//If one player has a king and the other has one piece they lose
	if wPieces == 1 {
		if bKings > 0 {
			return true, Black
		}
	}

	if bPieces == 1 {
		if wKings > 0 {
			return true, White
		}
	}

	return false, 0 //No winner
	//If a player has no playable moves they lose (checked in another part of the code)
}

func (board *BoardState) TryPromotion(){
	for x:=0;x<7;x++ {
		//Promote white y=0
		tile, _ := board.GetBoardTile(x,0)
		if tile.Team == White && tile.Full == Filled {
			tile.King = King
			board.SetBoardTile(x,0,tile)
		}

		//Promote black y=7
		tile, _ = board.GetBoardTile(x,7)
		if tile.Team == Black && tile.Full == Filled {
			tile.King = King
			board.SetBoardTile(x,7,tile)
		}
	}
}