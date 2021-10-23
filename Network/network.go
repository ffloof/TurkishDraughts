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

	results := make(chan struct{float64; board.BoardState})
	for _, branch := range options{
		go analyzeBranch(branch, myTeam, results) //TODO: implement proper worker pool instead of creating a thread for each initial move
	}

	var bestValue float64 
	var bestBoard board.BoardState

	for i := range options {
		result := <- results
		//fmt.Println(result.float64)
		//result.BoardState.Print()
		if i == 0 || (myTeam == board.White && result.float64 > bestValue) || (myTeam == board.Black && result.float64 < bestValue) {
			bestValue = result.float64
			bestBoard = result.BoardState
		}
	}

	fmt.Println(bestValue)
	return board.BoardToStr(&bestBoard)
}


func analyzeBranch (branch board.BoardState, myTeam board.Team, results chan struct{float64; board.BoardState}){
	if myTeam == board.White {
		results <- struct{float64; board.BoardState}{branch.BoardValue(9, -board.AlphaBetaMax, board.AlphaBetaMax, board.Black), branch}
	} else if myTeam == board.Black {
		results <- struct{float64; board.BoardState}{branch.BoardValue(9, -board.AlphaBetaMax, board.AlphaBetaMax, board.White), branch}
	}
}