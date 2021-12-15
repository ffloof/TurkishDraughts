package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmarkDepthsMM(t *testing.T, b board.BoardState, testDepth int32) {
	startTime := time.Now()
	board.Searches = 0

	value := b.MinMax(0, -999.0, 999.0, board.NewTable())

	duration := time.Since(startTime).Seconds()
	t.Log("---", testDepth)
	t.Log("Time:", float32(duration), "s")
	t.Log("Value:", value)
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")
}

func TestBenchNoAdvanced(t *testing.T){
	

	var i int32 = 5
	for i<=11 {
		//Config
		board.AdvanceWeight = 0.0
		board.TableDepthAllowedInaccuracy = 0
		board.MaxDepth = i
		board.MaximumHashDepth = 0

		b := board.CreateStartingBoard()
		benchmarkDepthsMM(t, b, i)
		i++
	}
}