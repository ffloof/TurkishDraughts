package board

type storedState struct {
	board BoardState //Stores board state since hash will have collisions
	value float32 //Stores previously evaluated value
	depth int32 //Stores depth to ensure that a position isnt too shallowly evaluated
}

type TransposTable struct {
	internal map[uint64]storedState
}

//Creates a new table and returns a pointer
func NewTable() *TransposTable {
	return &TransposTable{
		internal: make(map[uint64]storedState),
	}
}

//Returns if the board has been previously evaluated, and its value
func (table *TransposTable) Request(board *BoardState, depth int32) (bool, float32) {
	//Hash board state and load entry
	hash := board.hashBoard()
	entry, exists := table.internal[hash]

	if exists { //Check if it exists and if its the exact same board not just a collision 
		if entry.board == *board {
			//Checks if the board was evaluated at a depth greater than or within a certain range
			//Prevents using too shallowly evaluated moves
			if entry.depth - TableDepthAllowedInaccuracy <= depth { 
				return true, entry.value 
			}
		}
	}
	return false, 0.0 //Otherwise returns that it didn't find anything

}

//Sets the board
func (table *TransposTable) Set(board *BoardState, value float32, depth int32){
	//Hash board state and write to table
	hash := board.hashBoard()

	//Replace only if greater depth, ie more computationally expensive
	entry, exists := table.internal[hash]
	if !exists || depth <= entry.depth {
		//By saving deeper explored branches not only do we save the most time saving possiblity, we also perform far fewer writes increasing efficiency. 
		table.internal[hash] = storedState{*board, value, depth}
	}
}

//TODO: write unit test(s) for this
func (table *TransposTable) Turn(){
	for a,b := range table.internal {
		b.depth += 1
		if b.depth > MaximumHashDepth {
			delete(table.internal, a)
		} else {
			table.internal[a] = b
		}
	}
}

//63 bits of the hash are made of where pieces are (no other info about them) and 1 bit for whos turn it is
//This results in many collisions but this is necessary as no computer comes close to having enough memory to being able to address the whole 64bit range
//We only really care about storing a few expensive searches so this is ideal
func (board *BoardState) hashBoard() uint64 {
	return (board.Full << 1) | uint64(board.Turn)
}