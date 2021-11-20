package main

import (
	"testing"
	"time"
	"math/rand"
	"TurkishDraughts/Board"
)

func genRandomBoard() board.BoardState {
	return board.BoardState { 
		Turn: board.TileTeam(rand.Uint32()%2),
		Team: rand.Uint64(),
		King: rand.Uint64(), 
		Full: rand.Uint64(),
	}
}

func genRandomTile() board.Tile {
	return board.Tile { 
		Team: board.TileTeam(rand.Uint32()%2), 
		King: board.TileKing(rand.Uint32()%2), 
		Full: board.TileFull(rand.Uint32()%2),
	}
}

//Test GetBoardTile and SetBoardTile

func TestGetSetTile(t *testing.T){
	t.Log("Ready Get Set Go")
	rand.Seed(time.Now().UnixNano())
	randomBoard := genRandomBoard()

	for a:=0;a<100;a++{
		testTiles := []board.Tile{}

		for y:=-2;y<10;y++ {
			for x:=-2;x<10;x++ {
				randomTile := genRandomTile()
				testTiles = append(testTiles, randomTile)
				randomBoard.SetBoardTile(x,y,randomTile)
			}
		}

		i:=0
		for y:=-2;y<10;y++ {
			for x:=-2;x<10;x++ {
				myTile, valid := randomBoard.GetBoardTile(x,y)
				if (y > 7 || y < 0 || x > 7 || x < 0) {
					if valid {
						//Failed test
						t.Log("Position off board returned a value",x,y)
						t.Fail()
					}
				} else if !valid {
					//Failed test
					t.Log("Position on board failed to return a value", x, y)
					t.Fail()
				} else if testTiles[i] != myTile {
					//Failed test
					t.Log("Original tile does not match tile that was written/recieved from board", x, y)
					t.Fail()
				}
				i++
			}
		}
	}
}

//Test SwapTeam

func TestSwapTeam(t *testing.T){
	board1 := board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- --------")
	board1.SwapTeam()
	if board1.Turn != board.White {
		t.Log("Failed empty board swap test b->w")
		t.Fail()
	}
	board1.SwapTeam()
	if board1.Turn != board.Black {
		t.Log("Failed empty board swap test2 w->b")
	}

	board2 := board.BoardFromStr("wWbBWwBb wWbBWwBb bWbBWwBw WwbBWwbB wWbBWwBb wWbBWwBb wWbBWwBb wWbBWwBb")
	board2.SwapTeam()

	board3 := board.BoardFromStr("WWbBWWBb wWbBWwBb bWbBWwBw WwbBWwbB wWbBWwBb wWbBWwBb wWbBWwBb wWBBWwBB")
	board3.Turn = board.White

	if board2 != board3 {
		t.Log("Failed promotion test1")
		t.Log(board2.ToStr())
		t.Fail()
	}

	board2.SwapTeam()
	board3.Turn = board.Black

	if board2 != board3 {
		t.Log("Failed promotion test2")
		t.Log(board2.ToStr())
		t.Fail()
	}
}

//Test BoardToStr and BoardFromStr

func TestBoardStr(t *testing.T){
	board1 := board.BoardFromStr("wwwwWWWW bbbbBBBB -------- -------- -------- -------- -------- --------")
	board2 := board.BoardState { 
		Turn: 0, 
		Full: 0x000000000000FFFF,
		Team: 0x00000000000000FF,
		King: 0x000000000000F0F0,
	}

	if board1 != board2 {
		t.Log("Failed to make correct board from string")
		t.Log(board1)
		t.Fail()
	}

	if board2.ToStr() != "wwwwWWWW\nbbbbBBBB\n--------\n--------\n--------\n--------\n--------\n--------" {
		t.Log("Failed to make correct string from board")
		t.Log(board1.ToStr())
		t.Fail()
	}
}

//Test PlayerHasWon

func TestPlayerWin(t *testing.T){
	//Draw from no pieces
	board0 := board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- --------")
	win, winTeam, draw := board0.PlayerHasWon()
	if win || !draw { t.Fail() }

	//Win by no pieces remaining
	board1 := board.BoardFromStr("-------- -------- -------- -------- wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board1.PlayerHasWon()
	if !win || winTeam != board.White || draw { t.Fail()} 
	board2 := board.BoardFromStr("-------- -------- -------- -------- bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board2.PlayerHasWon()
	if !win || winTeam != board.Black || draw { t.Fail()} 

	//Not yet concluded
	board3 := board.BoardFromStr("-------- -------- b------- -------- wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board3.PlayerHasWon()
	if win || draw { t.Fail() }
	board4 := board.BoardFromStr("-------- -------- w------- -------- bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board4.PlayerHasWon()
	if win || draw { t.Fail() }

	//Won by king v 1
	board5 := board.BoardFromStr("-------- -------- b------- -------- Wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board5.PlayerHasWon()
	if !win || winTeam != board.White || draw { t.Fail()} 
	board6 := board.BoardFromStr("-------- -------- w------- -------- Bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board6.PlayerHasWon()
	if !win || winTeam != board.Black || draw { t.Fail()} 

	//Not yet concluded king v king + pieces
	board7 := board.BoardFromStr("-------- -------- B------- -------- Wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board7.PlayerHasWon()
	if win || draw { t.Fail() }
	board8 := board.BoardFromStr("-------- -------- W------- -------- Bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board8.PlayerHasWon()
	if win || draw { t.Fail() }

	//Draw king v king
	board9 := board.BoardFromStr("-------- -------- B------- -------- -------W -------- -------- --------")
	win, winTeam, draw = board9.PlayerHasWon()
	if win || !draw { t.Fail() }

	//Won 2king v king
	board10 := board.BoardFromStr("-------- B------- B------- -------- -------W -------- -------- --------")
	win, winTeam, draw = board10.PlayerHasWon()
	if !win || winTeam != board.Black || draw { t.Fail()} 
}

//Test table set and request

func TestTable(t *testing.T){
}

//Maybe test all moves board
//Maybe test max take boards
//Test AllMovesPawn
//Test AllMovesKing
//Test find king takes
//Test find pawn takes
//Maybe test random functions