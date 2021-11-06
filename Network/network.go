package network

import (
	"net/http"
	"fmt"
	"time"
	//"runtime"
	"runtime/debug"
	"TurkishDraughts/Board"
)

const netDepth = 10

type move struct {
	value float32 
	board board.BoardState
}

func Init(){
	http.HandleFunc("/", isAlive)
	http.HandleFunc("/black/", analyzeBlack)
	http.HandleFunc("/white/", analyzeWhite)

    http.ListenAndServe(":80", nil)
}

func isAlive(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Its alive!")
}

func analyzeBlack(w http.ResponseWriter, r *http.Request){
	b := board.BoardFromStr(r.URL.Path[7:])
	b.Turn = board.Black
	fmt.Fprintf(w, board.BoardToStr(Analyze(b,netDepth)))
}

func analyzeWhite(w http.ResponseWriter, r *http.Request){
	b := board.BoardFromStr(r.URL.Path[7:])
	b.Turn = board.White
	fmt.Fprintf(w, board.BoardToStr(Analyze(b,netDepth)))
}

func Analyze(b board.BoardState, depth uint32) *board.BoardState {
	if b.Turn == board.White {
		fmt.Println("================ WHITE ================")
	} else {
		fmt.Println("================ BLACK ================")
	}

	fmt.Println(time.Now().String())
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

	debug.FreeOSMemory()
	fmt.Println(time.Now().String())
	fmt.Println("Searches:", board.Searches/1000, "k")
	fmt.Println("Hits:", board.Hits/1000, "k")
	fmt.Println("Standing:", bestValue)

	return &bestBoard
}


func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan move, depth uint32) {
	branch.SwapTeam()
	output <- move {branch.MinMax(depth, -board.AlphaBetaMax, board.AlphaBetaMax, table), branch}
	debug.FreeOSMemory()
}