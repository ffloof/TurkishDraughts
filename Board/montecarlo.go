package board

import (
	"github.com/int8/gomcts"
)

type BoardAction BoardState

func (ba BoardAction) ApplyTo (gameState gomcts.GameState) gomcts.GameState {
	return BoardState(ba)
}

var IllegalBoards []BoardState = []BoardState{}

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

	plays := bs.ValidPlays()
	for _, prevB := range IllegalBoards {
		for i := range plays {
			if plays[i] == prevB {
				plays = remove(plays, i)
				break
			}
		}
	}

	if len(plays) == 0 {
		if bs.Turn == White {
			return gomcts.GameResult(-1), true
		}
		return gomcts.GameResult(1), true
	}

	return gomcts.GameResult(0), false
}

func (bs BoardState) GetLegalActions() []gomcts.Action {
	scuffedWorkaround := []gomcts.Action{}

	plays := bs.ValidPlays()
	for _, prevB := range IllegalBoards {
		for i := range plays {
			if plays[i] == prevB {
				plays = remove(plays, i)
				break
			}
		}
	}

	for _, v := range plays {
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
	return BoardState(choice.(BoardAction))
}

func remove(s []BoardState, i int) []BoardState {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}