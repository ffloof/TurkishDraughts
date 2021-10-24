package network

import (
	"net/http"
	"fmt"
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
	fmt.Fprintf(w, analyze(board.BoardFromStr(r.URL.Path[7:]), board.Black))
}

func analyzeWhite(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, analyze(board.BoardFromStr(r.URL.Path[7:]), board.White))
}

func analyze(b board.BoardState, myTeam board.Team) string {	
	options := b.MaxTakeBoards(myTeam)
	if len(options) == 0 {
		options = b.AllMoveBoards(myTeam)
	}

	var bestValue float64 
	var bestBoard board.BoardState

	for i, branch := range options{
		checkValue, checkBoard := analyzeBranch(branch, myTeam)
		if i == 0 || (myTeam == board.White && checkValue > bestValue) || (myTeam == board.Black && checkValue < bestValue) {
			bestValue = checkValue
			bestBoard = checkBoard
		}
	}


	fmt.Println(bestValue)
	return board.BoardToStr(&bestBoard)
}


func analyzeBranch (branch board.BoardState, myTeam board.Team) (float64, board.BoardState){
	if myTeam == board.White {
		return branch.BoardValue(9, -board.AlphaBetaMax, board.AlphaBetaMax, board.Black), branch
	} else if myTeam == board.Black {
		return branch.BoardValue(9, -board.AlphaBetaMax, board.AlphaBetaMax, board.White), branch
	}
	panic("Something went horribly wrong")
}