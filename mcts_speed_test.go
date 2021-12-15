package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmarkDepthsMCTS(t *testing.T, b board.BoardState, sims int) float64 {
	startTime := time.Now()
	board.Searches = 0

	board.MCTS(b, sims)

	duration := time.Since(startTime).Seconds()
	t.Log("---", sims)
	t.Log("Time:", float32(duration), "s")
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")

	return duration
}

func TestBenchMCTSVanilla(t *testing.T){
	//Config
	var i int = 128
	var lasttime float64
	for {
		lasttime = benchmarkDepthsMCTS(t, board.CreateStartingBoard(), i)
		if lasttime > 5.0 { break }
		i*=2
	}
	t.Log("===", i)
}