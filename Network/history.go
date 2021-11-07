package network

import (
	"TurkishDraughts/Board"
	"strings"
	"strconv"
	"fmt"
)

func ParseHistory(history string) board.BoardState {
	workingBoard := board.CreateStartingBoard()
	for _, moveStr := range strings.Fields(history) {
		if strings.Contains(moveStr, "x") { //Take
			startX, startY := parseCoordinate(moveStr[0:2])
			tile, _ := workingBoard.GetBoardTile(startX, startY)

			endX, endY := 0, 0
			for _, takeStr := range strings.Split(moveStr[3:], "x") {
				endX, endY = parseCoordinate(takeStr)
				for startX<endX {
					workingBoard.SetBoardTile(startX,startY,board.Tile{})
					startX++
				}
				for startY<endY {
					workingBoard.SetBoardTile(startX,startY,board.Tile{})
					startY++
				}
			}
			workingBoard.SetBoardTile(startX,startY,tile)
		} else { //Move
			startX, startY := parseCoordinate(moveStr[0:2]) 
			endX, endY := parseCoordinate(moveStr[3:5])
			tile, _ := workingBoard.GetBoardTile(startX, startY)
			workingBoard.SetBoardTile(endX,endY,tile)
			workingBoard.SetBoardTile(startX,startY,board.Tile{})
		}
		workingBoard.SwapTeam()
	}
	return workingBoard
}

func parseCoordinate(coordinate string) (int, int) {
	val, _ := strconv.ParseInt(coordinate,36,32)
	var y int = 7 - ((int(val) % 36) - 1)
	var x int = ((int(val) / 36) - 10)
	fmt.Println(x,y)
	return x,y
}