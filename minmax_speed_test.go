package main

import (
	"testing"
	"time"
	"runtime/debug"
	"TurkishDraughts/Board"
)

func benchmarkDepthsMM(t *testing.T, b board.BoardState, testDepth int32) float64 {
	startTime := time.Now()
	board.Searches = 0

	value := b.MinMax(0, -999.0, 999.0, board.NewTable())

	duration := time.Since(startTime).Seconds()
	t.Log("---", testDepth)
	t.Log("Time:", float32(duration), "s")
	t.Log("Value:", value)
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")

	//Forces memory to get cleared immediately so it will be ready for next test
	debug.FreeOSMemory()
	return duration
}



//Vanilla, default configuration
func TestBenchMMVanilla(t *testing.T){
	var i int32 = 5
	var lasttime float64
	for {
		//Config
		board.AdvanceWeight = 0.0
		board.TableDepthAllowedInaccuracy = 0
		board.MaxDepth = i
		board.MaximumHashDepth = i

		lasttime = benchmarkDepthsMM(t, board.CreateStartingBoard(), i)
		if lasttime > 15.0 { break }
		i++
	}
	t.Log("===", i)
}

//NoTable, doesn't use transposition table
func TestBenchMMNoTable(t *testing.T){
	var i int32 = 5
	var lasttime float64
	for i<=12 {
		//Config
		board.AdvanceWeight = 0.0
		board.TableDepthAllowedInaccuracy = 0
		board.MaxDepth = i
		board.MaximumHashDepth = 0

		lasttime = benchmarkDepthsMM(t, board.CreateStartingBoard(), i)
		if lasttime > 15.0 { break }
		i++
	}
	t.Log("===", i)
}


//FastTable, uses a table with settings slightly optimized
func TestBenchMMFastTable(t *testing.T){
	var i int32 = 5
	var lasttime float64
	for {
		//Config
		board.AdvanceWeight = 0.0
		board.TableDepthAllowedInaccuracy = 0
		board.MaxDepth = i
		board.MaximumHashDepth = i-2

		lasttime = benchmarkDepthsMM(t, board.CreateStartingBoard(), i)
		if lasttime > 15.0 { break }
		i++
	}
	t.Log("===", i)
}

//Advanced, default configuration + advance heuristic for evaluation
func TestBenchMMAdvanced(t *testing.T){
	var i int32 = 5
	var lasttime float64
	for {
		//Config
		board.AdvanceWeight = 0.1
		board.TableDepthAllowedInaccuracy = 0
		board.MaxDepth = i
		board.MaximumHashDepth = i

		lasttime = benchmarkDepthsMM(t, board.CreateStartingBoard(), i)
		if lasttime > 15.0 { break }
		i++
	}
	t.Log("===", i)
}

//CheatTable lets transposition table cheat slightly by looking at shallower depths
func TestBenchMMCheatTable(t *testing.T){
	var i int32 = 5
	var lasttime float64
	for {
		//Config
		board.AdvanceWeight = 0.0
		board.TableDepthAllowedInaccuracy = 2
		board.MaxDepth = i
		board.MaximumHashDepth = i - 2

		lasttime = benchmarkDepthsMM(t, board.CreateStartingBoard(), i)
		if lasttime > 15.0 { break }
		i++
	}
	t.Log("===", i)
}