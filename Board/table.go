package board

type storedState struct {
	board BoardState
	value float32
	depth int32
}

type TransposTable struct {
	internal map[uint64]storedState
}

func NewTable() *TransposTable {
	return &TransposTable{
		internal: make(map[uint64]storedState),
	}
}

func (table *TransposTable) request(board *BoardState, depth int32) (bool, float32) {
	//Hash board state and load entry
	hash := board.hashBoard()
	entry, exists := table.internal[hash]

	if exists {
		if entry.board == *board {
			if entry.depth - TableDepthAllowedInaccuracy <= depth { return true, entry.value }
		}
	}
	return false, 0.0

}

func (table *TransposTable) set(board *BoardState, value float32, depth int32){
	//Hash board state and write to table
	hash := board.hashBoard()

	//Replace only if greater depth
	entry, exists := table.internal[hash]
	if !exists || depth <= entry.depth { 
		//By saving shallower branches not only do we save the most time saving possiblity, we also perform far fewer writes increasing efficiency. 
		table.internal[hash] = storedState{*board, value, depth}
	}
}

func (board *BoardState) hashBoard() uint64 {
	return (board.Full << 1) | uint64(board.Turn)
}

/* func (board *BoardState) hashBoard2() uint64 {
	hash := (board.Full >> 16) & 0xFFFFFFFF
	hash = hash << 1
	hash |= uint64(board.Turn)
	return hash
} */