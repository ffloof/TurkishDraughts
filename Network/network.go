package network

import (
	"net/http"
	"fmt"
	"time"
	"TurkishDraughts/Board"
)

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

	for i, branch := range options{
		checkValue, checkBoard := analyzeBranch(branch, board.NewTable())
		if i == 0 || (b.Turn == board.White && checkValue > bestValue) || (b.Turn == board.Black && checkValue < bestValue) {
			bestValue = checkValue
			bestBoard = checkBoard
		}
	}

	fmt.Println("Searches:", board.Searches)
	fmt.Println(time.Now().String())

	fmt.Println(bestValue)
	return board.BoardToStr(&bestBoard)
}


func analyzeBranch (branch board.BoardState, table *board.TransposTable) (float64, board.BoardState){
	branch.SwapTeam()
	return branch.BoardValue(11, -board.AlphaBetaMax, board.AlphaBetaMax, table), branch
}