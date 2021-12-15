package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmarkDepthsMCTS(t *testing.T, b board.BoardState, sims int) {
	startTime := time.Now()
	board.Searches = 0

	board.MCTS(b, sims)

	duration := time.Since(startTime).Seconds()
	t.Log("---", sims)
	t.Log("Time:", float32(duration), "s")
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")
}

func TestBenchStartBoard(t *testing.T){
	//Config
	var i int = 64
	for i<=50000 {
		b := board.CreateStartingBoard()
		benchmarkDepthsMCTS(t, b, i)
		i*=2
	}	
}