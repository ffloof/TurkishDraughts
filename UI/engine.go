package ui

import (
	"fmt"
	"time"
	"runtime/debug"
	"TurkishDraughts/Board"
)

const netDepth = 10

type move struct {
	value float32 
	board board.BoardState
}

func Analyze(b board.BoardState, depth uint32) *board.BoardState {
	if b.Turn == board.White {
		fmt.Println("================ WHITE ================")
	} else {
		fmt.Println("================ BLACK ================")
	}

	startTime := time.Now()
	board.Searches = 0
	board.Hits = 0

	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	var bestValue float32 
	var bestBoard board.BoardState
	output := make(chan move)

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

	duration := time.Since(startTime).Seconds()
	fmt.Println("Time:", duration)
	fmt.Println("Searches:", board.Searches/1000, "k  Efficiency:", 100.0-(100.0*float64(board.Searches)/float64(board.Hits)) ,"%  Speed (k/s):",float64(board.Searches)/duration/1000.0,"k/s")
	fmt.Println("Standing:", bestValue)
	return &bestBoard
}


func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan move, depth uint32) {
	branch.SwapTeam()
	output <- move {branch.MinMax(depth, -board.AlphaBetaMax, board.AlphaBetaMax, table), branch}
	debug.FreeOSMemory()
}