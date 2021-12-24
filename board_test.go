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


	//Win by no pieces remaining
	board1 := board.BoardFromStr("-------- -------- -------- -------- wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw := board1.PlayerHasWon()
	if !win || winTeam != board.White || draw { 
		t.Log("1")
		t.Fail()
	} 
	board2 := board.BoardFromStr("-------- -------- -------- -------- bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board2.PlayerHasWon()
	if !win || winTeam != board.Black || draw {
		t.Log("2")
		t.Fail()
	} 

	//Not yet concluded
	board3 := board.BoardFromStr("-------- -------- b------- -------- wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board3.PlayerHasWon()
	if win || draw { 
		t.Log("3")
		t.Fail() 
	}
	board4 := board.BoardFromStr("-------- -------- w------- -------- bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board4.PlayerHasWon()
	if win || draw { 
		t.Log("4")
		t.Fail() 
	}

	//Won by king v 1
	board5 := board.BoardFromStr("-------- -------- b------- -------- Wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board5.PlayerHasWon()
	if !win || winTeam != board.White || draw {
		t.Log("5") 
		t.Fail()
	} 
	board6 := board.BoardFromStr("-------- -------- w------- -------- Bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board6.PlayerHasWon()
	if !win || winTeam != board.Black || draw { 
		t.Log("6")
		t.Fail()
	} 

	//Not yet concluded king v king + pieces
	board7 := board.BoardFromStr("-------- -------- B------- -------- Wwwwwwww wwwwwwww -------- --------")
	win, winTeam, draw = board7.PlayerHasWon()
	if win || draw { 
		t.Log("7")
		t.Fail() 
	}
	board8 := board.BoardFromStr("-------- -------- W------- -------- Bbbbbbbb bbbbbbbb -------- --------")
	win, winTeam, draw = board8.PlayerHasWon()
	if win || draw { 
		t.Log("8")
		t.Fail() 
	}

	//Draw king v king
	board9 := board.BoardFromStr("-------- -------- B------- -------- -------W -------- -------- --------")
	win, winTeam, draw = board9.PlayerHasWon()
	if win || !draw { 
		t.Log("9")
		t.Fail() 
	}
}

//Test table set and request

func TestTable(t *testing.T){
	defaultBoard := board.CreateStartingBoard()
	testTable := board.NewTable(100, 3)

	testTable.Set(defaultBoard, 1.0, 6)
	exists, value := testTable.Request(defaultBoard, 10)
	if !exists || value != 1.0 {
		t.Log("Table write/read failed")
		t.Fail()
	}

	testTable.Set(defaultBoard, 2.0, 7)
	exists, value = testTable.Request(defaultBoard, 10)
	if !exists || value != 1.0 {
		t.Log("Table lower depth write")
		t.Fail()
	}

	testTable.Set(defaultBoard, 3.0, 5)
	exists, value = testTable.Request(defaultBoard, 10)
	if !exists || value != 3.0 {
		t.Log("Table higher depth write failed")
		t.Fail()
	}

	exists, value = testTable.Request(defaultBoard, 1)
	if exists {
		t.Log("Read out of bounds of allowed inaccuracy")
		t.Fail()
	}

	exists, value = testTable.Request(defaultBoard, 2)
	if !exists {
		t.Log("Failed to read on edge of bounds of allowed inaccuracy")
		t.Fail()
	}

	exists, value = testTable.Request(defaultBoard, 3)
	if !exists {
		t.Log("Failed to read within bounds of allowed inaccuracy")
		t.Fail()
	}

	for i:=100;i>0;i-- {
		randomBoard := genRandomBoard()
		testTable.Set(randomBoard, 8.0, int32(i))
		exists, value = testTable.Request(randomBoard, int32(i))
		if !exists || value != 8.0 {
			t.Log("Failed random read write test")
			t.Fail()
		}
	}
}

//Tests if all of b are in a, does not care about duplicates in a
func equalUnsorted(a []board.BoardState, b []board.BoardState) bool {
	a = board.RemoveDuplicateValues(a)
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

//Test AllMovesPawn
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
	board1 := board.BoardFromStr("-------- -------- -------- ---W---- -------- -------- -------- --------")

	passed := equalUnsorted(board1.AllMovesKing(3,3), []board.BoardState{
		board.BoardFromStr("-------- -------- -------- -------- ---W---- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -------- -------- ---W---- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -------- -------- -------- ---W---- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- ---W----"),                                                                                          
		board.BoardFromStr("-------- -------- ---W---- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- ---W---- -------- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("---W---- -------- -------- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- --W----- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -W------ -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- W------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- ----W--- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -----W-- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- ------W- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("-------- -------- -------- -------W -------- -------- -------- --------"),
	})

	if !passed {
		t.Log("King free move test failed")
		t.Fail()
	}

	
	//Blocked move by piece on sides at close and long range and by a wall
	board2 := board.BoardFromStr("--bW--b- -------- -------- -------- -------- -------- -------- --------")

	passed = equalUnsorted(board2.AllMovesKing(3,0), []board.BoardState{
		board.BoardFromStr("--b---b- ---W---- -------- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- ---W---- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- -------- ---W---- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- -------- -------- ---W---- -------- -------- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- -------- -------- -------- ---W---- -------- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- -------- -------- -------- -------- ---W---- --------"),                                                                                          
		board.BoardFromStr("--b---b- -------- -------- -------- -------- -------- -------- ---W----"),                                                                                          
		board.BoardFromStr("--b-W-b- -------- -------- -------- -------- -------- -------- --------"),                                                                                          
		board.BoardFromStr("--b--Wb- -------- -------- -------- -------- -------- -------- --------"),
	})

	if !passed {
		t.Log("King blocked move test failed")
		t.Fail()
	}
}

//Test find king takes
func TestKingTakes(t *testing.T){
	//Test none, against wall, next to, spaced out
	board1 := board.BoardFromStr("-------- -------- -------- -b------ bW---b-- -------- -------- -B------")

	_, compare, _ := board1.FindKingTakes(1,4,0,[2]int{0,0})
	
	passed := equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- -------- -W------ -------- b----b-- -------- -------- -B------"),                                                                                                                   
		board.BoardFromStr("-------- -W------ -------- -------- b----b-- -------- -------- -B------"),                                                                                                                   
		board.BoardFromStr("-W------ -------- -------- -------- b----b-- -------- -------- -B------"),                                                                                                                   
		board.BoardFromStr("-------- -------- -------- -b------ b-----W- -------- -------- -B------"),                                                                                                                   
		board.BoardFromStr("-------- -------- -------- -b------ b------W -------- -------- -B------"),
	})

	if !passed {
		t.Log("Failed king take collision tests")
		t.Fail()
	}

	//Test multi take in all directions even over previous removed pieces
	board2 := board.BoardFromStr("-----b-- ------b- ---bb--- -------- ----W--- -------- -------- --------")

	_, compare, _ = board2.FindKingTakes(4,4,0,[2]int{0,0})

	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- -------- --W----- -------- -------- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- -W------ -------- -------- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- W------- -------- -------- -------- -------- --------"),  
	})

	if !passed {
		t.Log("Failed piece king take removal test")
		t.Fail()
	}

	//Test branching situation with 1 correct max
	board3 := board.BoardFromStr("-------- b------b ----W--- -------- ----b--- b------- ----b--- --b---B-")

	_, compare, _ = board3.FindKingTakes(4,2,0,[2]int{0,0})
	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("W------- -------b -------- -------- -------- -------- -------- ------B-"),
	})

	if !passed {
		t.Log("Failed branch king take 1 max test")
		t.Fail()
	}

	//Test branching situation with 2 correct max
	board4 := board.BoardFromStr("-------- b------b ----W--- -------b ----b--- b------- ----b--- --b---B-")

	_, compare, _ = board4.FindKingTakes(4,2,0,[2]int{0,0})
	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("W------- -------b -------- -------b -------- -------- -------- ------B-"),                                                                                                                   
		board.BoardFromStr("-------W b------- -------- -------- -------- b------- -------- --b-----"),
	})

	if !passed {
		t.Log("Failed branch king take 2 max test")
		t.Fail()
	}

	//Test taking backwards
	board5 := board.BoardFromStr("-------- -------- -------- --b-Wb-- -------- -------- -------- --------")

	_, compare, _ = board5.FindKingTakes(4,3,0,[2]int{0,0})
	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- -------- -------- -W---b-- -------- -------- -------- --------"),                                                                                                        
		board.BoardFromStr("-------- -------- -------- W----b-- -------- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- -------- --b---W- -------- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- -------- --b----W -------- -------- -------- --------"),
	})

	if !passed {
		t.Log("Failed backwards king take test")
		t.Fail()
	}

	//Jump over multiple check
	board6 := board.BoardFromStr("-------- ---b---- ---b---- -bbW-bb- -------- ---b---- ---b---- --------")

	_, compare, _ = board6.FindKingTakes(3,3,0,[2]int{0,0})

	if len(compare) != 1 {
		t.Log("Failed king multi piece take")
		t.Fail()
	}

	//Friendly fire check
	board7 := board.BoardFromStr("-------- ---w---- ---b---- --WW-bW- -------- ---w---- -------- --------")

	_, compare, _ = board7.FindKingTakes(3,3,0,[2]int{0,0})

	if len(compare) != 1 || board7 != compare[0] {
		t.Log("Failed king take friendly fire test")
		t.Fail()
	}
}

//Test find pawn takes
func TestPawnTakes(t *testing.T){
	//Test normal takes excluding backwars both colors
	board1 := board.BoardFromStr("-------- -------- ---b---- --bwb--- ---b---- -------- -------- --------")
	_, compare, _ := board1.FindPawnTakes(3,3,0)

	passed := equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- ---w---- -------- --b-b--- ---b---- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- ---b---- -w--b--- ---b---- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- ---b---- --b--w-- ---b---- -------- -------- --------"), 
	})

	if !passed {
		t.Log("Failed white pawn take test")
		t.Fail()
	}

	board2 := board.BoardFromStr("-------- -------- ---w---- --wbw--- ---w---- -------- -------- --------")
	_, compare, _ = board2.FindPawnTakes(3,3,0)

	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- -------- ---w---- --w-w--- -------- ---b---- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- ---w---- -b--w--- ---w---- -------- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- -------- ---w---- --w--b-- ---w---- -------- -------- --------"), 
	})

	if !passed {
		t.Log("Failed black pawn take test")
		t.Fail()
	}

	//Test take against wall
	board3 := board.BoardFromStr("-b------ bw------ -------- -------- -------- -------- -------- --------")
	_, compare, _ = board3.FindPawnTakes(1,1,0)
	if len(compare) != 1 || board3 != compare[0] {
		t.Log("Failed pawn take wall test")
		t.Fail()
	}


	//Test branching situation with 1 correct max
	board4 := board.BoardFromStr("-------- -------- --b----- -------- --b---b- ---b-b-- ----b--- --wb----")
	_, compare, _ = board4.FindPawnTakes(2,7,0)

	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- --w----- -------- -------- ------b- -----b-- -------- --------"),
	})

	if !passed {
		t.Log("Failed branching pawn take test 1")
		t.Fail()
	}

	//Test branching situation with 2 correct max
	board5 := board.BoardFromStr("-------- -------- --b---b- -------- --b---b- ---b-b-- ----b--- --wb----")
	_, compare, _ = board5.FindPawnTakes(2,7,0)

	passed = equalUnsorted(compare, []board.BoardState{
		board.BoardFromStr("-------- --w----- ------b- -------- ------b- -----b-- -------- --------"),                                                                                                                   
		board.BoardFromStr("-------- ------w- --b----- -------- --b----- ---b---- -------- --------"),
	})

	if !passed {
		t.Log("Failed branching pawn take test 2")
		t.Fail()
	}

	//Test against friendly fire / multiple pieces
	board6 := board.BoardFromStr("-------- ---w---- -wbwW--- -------- -------- -------- -------- --------")
	_, compare, _ = board6.FindPawnTakes(3,2,0)
	if len(compare) != 1 || board6 != compare[0] {
		t.Log("Failed pawn take friendly fire test")
		t.Fail()
	}
}

//Maybe test all moves board
//Maybe test max take boards