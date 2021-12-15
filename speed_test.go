package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmarkDepths(t *testing.T, b board.BoardState, sims int) {
	startTime := time.Now()
	board.Hits = 0
	board.Searches = 0

	board.MCTS(b, sims)

	duration := time.Since(startTime).Seconds()
	t.Log("---", sims)
	t.Log(int(float64(sims)/duration),"/s")
	t.Log("Time:", float32(duration), "s")
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")
}

func TestBenchStartBoard(t *testing.T){
	//Config
	var i int = 64
	for i<=50000 {
		b := board.CreateStartingBoard()
		benchmarkDepths(t, b, i)
		i*=2
	}	
}