package board

import (
	"github.com/int8/gomcts"
)

var ( //Values used for seeing benchmark performance
	Hits = 0
	Searches = 0
)

type BoardAction BoardState

func (ba BoardAction) ApplyTo (gameState gomcts.GameState) gomcts.GameState {
	return BoardState(ba)
}

//Implement game state interface
func (bs BoardState) EvaluateGame() (gomcts.GameResult, bool) {
	playerWon, winWhite, playerDrew := bs.PlayerHasWon()
	
	if playerWon {
		if winWhite == White {
			return gomcts.GameResult(1), true
		} 
		return gomcts.GameResult(-1), true
	} else if playerDrew {
		return gomcts.GameResult(0), true
	}

	if len(bs.ValidPlays()) == 0 {
		if bs.Turn == White {
			return gomcts.GameResult(-1), true
		}
		return gomcts.GameResult(1), true
	}

	return gomcts.GameResult(0), false
}

func (bs BoardState) GetLegalActions() []gomcts.Action {
	scuffedWorkaround := []gomcts.Action{}
	for _, v := range bs.ValidPlays() {
		v.SwapTeam()
		scuffedWorkaround = append(scuffedWorkaround, BoardAction(v)) 
	}
	return scuffedWorkaround
}

func (bs BoardState) IsGameEnded() bool {
	playerWon, _, playerDrew := bs.PlayerHasWon()
	
	if playerWon || playerDrew {
		return true 
	}

	if len(bs.ValidPlays()) == 0 {
		return true
	}

	return false
}



func (bs BoardState) NextToMove() int8 {
	if bs.Turn == White {
		return 1
	} else {
		return -1
	}
}


func MCTS(b BoardState, nodes int) BoardState {
	choice := gomcts.MonteCarloTreeSearch(b, gomcts.DefaultRolloutPolicy, nodes)
	x := choice.(BoardAction)
	y := BoardState(x)
	//y.Print()
	return y
}
