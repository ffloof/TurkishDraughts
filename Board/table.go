package board

import (
)

type storedState struct {
	board BoardState
	value float64
	depth uint32
}

type TransposTable struct {
	internal map[uint64]storedState
}

func NewTable() *TransposTable {
	return &TransposTable{
		internal: make(map[uint64]storedState),
	}
}

func (table *TransposTable) Request(board *BoardState) (bool, float64) {
	//Hash board state and load entry
	hash := board.hashBoard()
	entry, exists := table.internal[hash]

	if exists {
		if entry.board == *board {
			return true, entry.value
		}
	}
	return false, 0.0

}

func (table *TransposTable) Set(board *BoardState, value float64, depth uint32){
	//Hash board state and write to table
	hash := board.hashBoard()

	//Replace only if greater depth
	entry, exists := table.internal[hash]
	if !exists || depth >= entry.depth {
		table.internal[hash] = storedState{*board, value, depth}
	}
}

func (board *BoardState) hashBoard() uint64 {
	return (board.Full << 1) | uint64(board.Turn)
}

func (board *BoardState) hashBoard2() uint64 {
	hash := ((board.Full >> 16) << 32)
	hash |= uint64(board.Turn)
	return hash
}