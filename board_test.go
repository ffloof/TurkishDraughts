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

}

//Test PlayerHasWon

func TestPlayerWin(t *testing.T){

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