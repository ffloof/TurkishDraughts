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

//Uint64 acts as an array of bits, 8x8 grid = 64 bits
type BoardState struct { 
	Turn TileTeam
	Team uint64 
	King uint64 
	Full uint64 
}

//Returns a board tile and a boolean of if it exists
func (bs *BoardState) GetBoardTile(xed int, yed int) (Tile, bool) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 { //Check if its on the board
		x := uint64(xed)
		y := uint64(yed)

		//Use bitwise operations to extract specific tile
		team := (bs.Team >> (y*8+x)) & 1
		king := (bs.King >> (y*8+x)) & 1
		full := (bs.Full >> (y*8+x)) & 1

		return Tile{TileTeam(team),TileKing(king),TileFull(full)}, true
	}
	return Tile{}, false //Return empty tile and that it isnt on board
}

//Sets a tile at a given position
func (bs *BoardState) SetBoardTile(xed int, yed int, t Tile) {
	if -1 < xed && xed < 8 && -1 < yed && yed < 8 { //Checks if the position is on the board
		x := uint64(xed)
		y := uint64(yed)

		//Bitwise logic to set tile at an index
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

//Swaps the BoardState.Turn value and promotes pawns at end of turn
func (bs *BoardState) SwapTeam(){
	bs.tryPromotion()
	if bs.Turn == White {
		bs.Turn = Black
	} else {
		bs.Turn = White
	}
}

//Converts board into a string
func (bs *BoardState) ToStr() string {
	fullStr := ""
	for y:=0;y<8;y++ {
		//Each line is a string of pieces in that row
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

//Prints the board for debugging
func (bs *BoardState) Print(){
	for _, line := range strings.Fields(bs.ToStr()) {
		fmt.Println(line)
	} 
}

//Returns starting board
func CreateStartingBoard() BoardState {
	b := BoardFromStr("-------- bbbbbbbb bbbbbbbb -------- -------- wwwwwwww wwwwwwww --------")
	b.Turn = White
	return b
}

//Converts from string to board for debugging/testing
func BoardFromStr(str string) BoardState {
	rows := strings.Fields(str)
	var board BoardState

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if string(rows[y][x]) == "-" { 
				board.SetBoardTile(x,y,Tile{}) //empty tile
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

//Returns if a board IsWon, TeamWon, IsDraw
func (bs *BoardState) PlayerHasWon() (bool, TileTeam, bool) { 
	//Sum all the pieces on the board
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

	//If either player is out of pieces they lose
	if wPieces == 0 {
		return true, Black, false
	} 

	if bPieces == 0 {
		return true, White, false
	}

	//If one player has at least 1 king and the other has one piece they lose
	if wPieces == 1 && bPieces == 1 {
		if wKings > 0 {
			if bKings > 0 {
				return false, 0, true //Draw
			}
			return true, White, false //White wins
		}

		if bKings > 0 {
			return true, Black, false //Black wins
		}
	}

	return false, 0.0, false //No winner or draw
	//If a player has no playable moves they lose (checked in another part of the code)
}

//Searches end of each file and promotes pawns to kings if either side has reached their respective end
func (board *BoardState) tryPromotion(){
	for x:=0;x<8;x++ {
		//Promote white at y=0 (top board)
		tile, _ := board.GetBoardTile(x,0)
		if tile.Team == White && tile.Full == Filled {
			tile.King = King
			board.SetBoardTile(x,0,tile)
		}

		//Promote black at y=7 (bottom board)
		tile, _ = board.GetBoardTile(x,7)
		if tile.Team == Black && tile.Full == Filled {
			tile.King = King
			board.SetBoardTile(x,7,tile)
		}
	}
}