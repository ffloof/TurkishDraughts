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
	defaultBoard := board.CreateStartingBoard()
	testTable := board.NewTable()

	board.TableDepthAllowedInaccuracy = 3
	board.MaximumHashDepth = 100

	testTable.Set(&defaultBoard, 1.0, 6)
	exists, value := testTable.Request(&defaultBoard, 10)
	if !exists || value != 1.0 {
		t.Log("Table write/read failed")
		t.Fail()
	}

	testTable.Set(&defaultBoard, 2.0, 7)
	exists, value = testTable.Request(&defaultBoard, 10)
	if !exists || value != 1.0 {
		t.Log("Table lower depth write")
		t.Fail()
	}

	testTable.Set(&defaultBoard, 3.0, 5)
	exists, value = testTable.Request(&defaultBoard, 10)
	if !exists || value != 3.0 {
		t.Log("Table higher depth write failed")
		t.Fail()
	}

	exists, value = testTable.Request(&defaultBoard, 1)
	if exists {
		t.Log("Read out of bounds of allowed inaccuracy")
		t.Fail()
	}

	exists, value = testTable.Request(&defaultBoard, 2)
	if !exists {
		t.Log("Failed to read on edge of bounds of allowed inaccuracy")
		t.Fail()
	}

	exists, value = testTable.Request(&defaultBoard, 3)
	if !exists {
		t.Log("Failed to read within bounds of allowed inaccuracy")
		t.Fail()
	}

	for i:=100;i>0;i-- {
		randomBoard := genRandomBoard()
		testTable.Set(&randomBoard, 8.0, int32(i))
		exists, value = testTable.Request(&randomBoard, int32(i))
		if !exists || value != 8.0 {
			t.Log("Failed random read write test")
			t.Fail()
		}
	}
}

//Maybe test all moves board
//Maybe test max take boards
//Test AllMovesPawn

func equalUnsorted(a []board.BoardState, b []board.BoardState) bool {
	if len(a) != len(b) { return false }
	for _, c := range b {
		included := false
		for _, d := range a {
			if c == d { 
				included = true
				break
			}
		}
		if !included { return false }
	}
	return true
}

func TestPawnMoves(t *testing.T){
	//Free move both sides
	board1 := board.BoardFromStr("-------- B------- -------- ---w---- -------- -------- -----b-- -bbb----")
	
	passed := equalUnsorted(board1.AllMovesPawn(3,3), []board.BoardState{
		board.BoardFromStr("-------- B------- ---w---- -------- -------- -------- -----b-- -bbb----"),                                                                                       
		board.BoardFromStr("-------- B------- -------- --w----- -------- -------- -----b-- -bbb----"),                                                                                         
		board.BoardFromStr("-------- B------- -------- ----w--- -------- -------- -----b-- -bbb----"),
	})
	
	if !passed {
		t.Log("Failed white pawn move test")
		t.Fail()
	}

	board2 := board.BoardFromStr("-------- W------- -------- ---b---- -------- -------- -----w-- -www----")
	
	passed = equalUnsorted(board2.AllMovesPawn(3,3), []board.BoardState{
		board.BoardFromStr("-------- W------- -------- -------- ---b---- -------- -----w-- -www----"),
		board.BoardFromStr("-------- W------- -------- --b----- -------- -------- -----w-- -www----"),
		board.BoardFromStr("-------- W------- -------- ----b--- -------- -------- -----w-- -www----"),
	})

	if !passed {
		t.Log("Failed black pawn move test")
		t.Fail()
	}

	//Blocked move by pieces on some sides
	board3 := board.BoardFromStr("---b---- --bw---- -------- -------- -------- -------- -------- --------")
	passed = equalUnsorted(board3.AllMovesPawn(3,1), []board.BoardState{
		board.BoardFromStr("---b---- --b-w--- -------- -------- -------- -------- -------- --------"),	
	})

	if !passed {
		t.Log("Failed blocked move test")
		t.Fail()
	}


	//Blocked move by wall
	board4 := board.BoardFromStr("-------- w------- -------- -------- -------- -------- -------- --------")
	passed = equalUnsorted(board4.AllMovesPawn(0,1), []board.BoardState{
		board.BoardFromStr("w------- -------- -------- -------- -------- -------- -------- --------"),
		board.BoardFromStr("-------- -w------ -------- -------- -------- -------- -------- --------"),
	})

	if !passed {
		t.Log("Failed wall move test")
		t.Fail()
	}
}

//Test AllMovesKing
func TestKingMoves(t *testing.T){
	//Free move all sides
	//board1 := board.BoardFromStr("-------- -------- -------- ---W---- -------- -------- -------- --------")

	/*
	-------- -------- -------- -------- ---W---- -------- -------- --------                                                                                          
-------- -------- -------- -------- -------- ---W---- -------- --------                                                                                          
-------- -------- -------- -------- -------- -------- ---W---- --------                                                                                          
-------- -------- -------- -------- -------- -------- -------- ---W----                                                                                          
-------- -------- ---W---- -------- -------- -------- -------- --------                                                                                          
-------- ---W---- -------- -------- -------- -------- -------- --------                                                                                          
---W---- -------- -------- -------- -------- -------- -------- --------                                                                                          
-------- -------- -------- --W----- -------- -------- -------- --------                                                                                          
-------- -------- -------- -W------ -------- -------- -------- --------                                                                                          
-------- -------- -------- W------- -------- -------- -------- --------                                                                                          
-------- -------- -------- ----W--- -------- -------- -------- --------                                                                                          
-------- -------- -------- -----W-- -------- -------- -------- --------                                                                                          
-------- -------- -------- ------W- -------- -------- -------- --------                                                                                          
-------- -------- -------- -------W -------- -------- -------- --------
*/
	//Blocked move by piece on sides at close and long range and by a wall
	//board2 := board.BoardFromStr("--bW--b- -------- -------- -------- -------- -------- -------- --------")
/*
--b---b- ---W---- -------- -------- -------- -------- -------- --------                                                                                          
--b---b- -------- ---W---- -------- -------- -------- -------- --------                                                                                          
--b---b- -------- -------- ---W---- -------- -------- -------- --------                                                                                          
--b---b- -------- -------- -------- ---W---- -------- -------- --------                                                                                          
--b---b- -------- -------- -------- -------- ---W---- -------- --------                                                                                          
--b---b- -------- -------- -------- -------- -------- ---W---- --------                                                                                          
--b---b- -------- -------- -------- -------- -------- -------- ---W----                                                                                          
--b-W-b- -------- -------- -------- -------- -------- -------- --------                                                                                          
--b--Wb- -------- -------- -------- -------- -------- -------- --------

*/
}

//Test find king takes
//Test find pawn takes
//Maybe test random functions