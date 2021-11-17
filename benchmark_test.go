package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmark(t *testing.T, b *board.BoardState, testDepth int32) {
	startTime := time.Now()
	board.Searches = 0
	board.Hits = 0
	board.Depth = testDepth

	value := b.MinMax(testDepth, -board.AlphaBetaMax, board.AlphaBetaMax, board.NewTable())

	duration := time.Since(startTime).Seconds()
	t.Log("---", testDepth)
	t.Log("Time:", duration)
	t.Log("Searches:", board.Searches/1000, "k")
	t.Log("Efficiency:", float32(100.0-(100.0*float64(board.Searches)/float64(board.Hits))) ,"%")
	t.Log("Speed(k/s):",float32(float64(board.Searches)/duration/1000.0),"k/s")
	t.Log("Standing:", value)
}

func TestBenchDefaultBoard(t *testing.T){
	var i int32 = 5
	for i<12 {
		b := board.CreateStartingBoard()
		benchmark(t, &b, i)
		i++
	}
	
}

/*

	startTime := time.Now()
	board.Searches = 0
	board.Hits = 0

	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	var bestValue float32 
	var bestBoard board.BoardState
	output := make(chan PossibleMove)

	for _, branch := range options{
		go analyzeBranch(branch, board.NewTable(), output, depth)
	}

	for i := range options {
		check := <- output
		checkValue := check.value
		checkBoard := check.board
		if i == 0 || (b.Turn == board.White && checkValue > bestValue) || (b.Turn == board.Black && checkValue < bestValue) {
			bestValue = checkValue
			bestBoard = checkBoard
		}

		fmt.Println(i+1, "/", len(options), "=", checkValue)
	}

	
	return &bestBoard


*/