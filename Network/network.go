package network

import (
	"net/http"
	"fmt"
	"time"
	//"runtime"
	"runtime/debug"
	"TurkishDraughts/Board"
)

const Depth = 13

type move struct {
	value float64 
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
	fmt.Fprintf(w, analyze(b))
}

func analyzeWhite(w http.ResponseWriter, r *http.Request){
	b := board.BoardFromStr(r.URL.Path[7:])
	b.Turn = board.White
	fmt.Fprintf(w, analyze(b))
}

func analyze(b board.BoardState) string {
	fmt.Println(time.Now().String())
	board.Searches = 0

	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	var bestValue float64 
	var bestBoard board.BoardState
	//TODO: find a better solution that RWMutex, it really slows it down, might just make a table for each thread
	output := make(chan move)

	for _, branch := range options{
		go analyzeBranch(branch, board.NewTable(), output)
	}

	for i := range options {
		check := <- output
		checkValue := check.value
		checkBoard := check.board
		if i == 0 || (b.Turn == board.White && checkValue > bestValue) || (b.Turn == board.Black && checkValue < bestValue) {
			bestValue = checkValue
			bestBoard = checkBoard
		}

		fmt.Println(i+1, "/", len(options))
	}

	debug.FreeOSMemory()
	fmt.Println(time.Now().String())
	fmt.Println("Searches:", board.Searches)
	fmt.Println("Standing:", bestValue)

	return board.BoardToStr(&bestBoard)
}


func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan move) {
	branch.SwapTeam()
	output <- move {branch.BoardValue(Depth, -board.AlphaBetaMax, board.AlphaBetaMax, table), branch}
	debug.FreeOSMemory()
}