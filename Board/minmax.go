package board

const (
	//Weights for pieces/wins
	//Win should always be greater than the theoretical max of value of a board where one side gets 16 kings
	alphaBetaMax float32 = 999.0
	winWeight float32 = 100.0
	kingWeight float32 = 5.0
	pawnWeight float32 = 1.0
)

var (
	//What depth minmax searches to
	MaxDepth int32 = 10 //default 10
	
	//Heuristic weight of how far advanced a sides pawn pieces are from promotion //TODO: make it increase as piece count decreases?
	//When its not 0.0 it makes ab pruning much slower put pushes the engine to play better in the long term
	AdvanceWeight float32 = 0.0 //default 0.0 or 0.1
)

//Evaluates recursively the value of a board using the minmax algorithm
//Board value is always when in whites favor positive and blacks favor negative
func (bs BoardState) MinMax(depth int32, alpha, beta float32, table *TransposTable) float32 {
	//Checks table to see if theres already an entry for this board
	
	if alreadyChecked, prevValue := table.Request(bs, depth); alreadyChecked {
		return prevValue //If there is one no need to research this branch of the tree
	}
	

	//Checks if the board is won or if players have drawed and returns accordingly
	playerWon, winWhite, playerDrew := bs.PlayerHasWon()
	
	if playerWon {
		if winWhite == White {
			return winWeight
		} 
		return -winWeight
	} else if playerDrew {
		return 0.0
	}

	if depth == MaxDepth {
		return bs.RawBoardValue(AdvanceWeight)
	}	

	options := bs.ValidPlays()
	for _, prevB := range IllegalBoards {
		for i := range options {
			if options[i] == prevB {
				options = remove(options, i)
				break
			}
		}

	}
	if len(options) == 0 { //If a payer has no legal moves they lose
		if bs.Turn == White { 
			return -winWeight
		} else { 
			return winWeight
		}
	}
	
	var value float32

	//Search for the best possible value move
	if bs.Turn == White {
		value = -alphaBetaMax
		for _, branch := range options { //Search each possible move with minmax

			v := branch.MinMax(depth+1, alpha, beta, table)
			
			//AB pruning to speed up tree search
			if v > value {
				value = v
			}

			if value >= beta { break }
			if value > alpha { alpha = value }
		}
	} else { //Same just from black's perspective
		value = alphaBetaMax
		for _, branch := range options {

			v := branch.MinMax(depth+1, alpha, beta, table)

			if v < value {
				value = v
			}

			if value <= alpha { break }
			if value < beta { beta = value }
		}
	}
	
	//Cache move in the table to speed up later searches of identical situations
	table.Set(bs, value, depth)

	//Return final best value of a move found from this board
	return value
}

//Gets the value of the board by summing piece weights and how far advanced a sides pieces are
func (bs *BoardState) RawBoardValue(advanceWeight float32) float32 { 
	wPawns := 0
	wKings := 0
	bPawns := 0
	bKings := 0
	netAdvance := 0

	//Sum amount of pieces and how advanced they are up the board
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			piece, _ := bs.GetBoardTile(x,y)
			if piece.Full == Empty { continue }
			if piece.Team == White {
				if piece.King == King {
					wKings += 1
				} else {
					wPawns += 1
					netAdvance += (7-y)
				}
			} else {
				if piece.King == King {
					bKings += 1
				} else {
					bPawns += 1
					netAdvance -= (y)
				}
			}
		}
	}

	//Some arithmetic to get the final board value
	var value float32 = 0.0

	value += pawnWeight * float32(wPawns - bPawns)
	value += kingWeight * float32(wKings - bKings)
	value += advanceWeight * float32(netAdvance)

	return value
}