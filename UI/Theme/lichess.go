package theme

//Theme based off the chess board/pieces at lichess.org, no affiliation.

import (
	"TurkishDraughts/Board"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)


type LichessTheme struct {} 

var lichessDimensions = dimensions {
	Height: 900.0,
	Border: 20.0,
	Gaps: 0.0,
}

func (t LichessTheme) GetMouseData(win *pixelgl.Window) (bool, bool, int) {
	return lichessDimensions.getMouseData(win)
}

func (t LichessTheme) DrawBoard(imd *imdraw.IMDraw){
	imd.Color = color.RGBA{0xED, 0xEB, 0xE9, 0xFF} //Board background color
	imd.Push(pixel.V(0.0,0.0), pixel.V(lichessDimensions.Height, lichessDimensions.Height))
	imd.Rectangle(0.0)

	for i:=0;i<64;i++ {
		if (i/8 + i%8)%2 == 0 { //Checkerboard effect (x+y)%2
			imd.Color = color.RGBA{0xB5, 0x88, 0x63, 0xFF} //Tile Dark Color
		} else {
			imd.Color = color.RGBA{0xF0, 0xD9, 0xB5, 0xFF} //Tile Light Color
		}
		x1, y1 := lichessDimensions.getTilePos(i,0.0,0.0)
		x2, y2 := lichessDimensions.getTilePos(i,1.0,1.0)
		imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
		imd.Rectangle(0.0)
	}
}

func (t LichessTheme) DrawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	size := lichessDimensions.getTileSpace()

	whiteColor := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF} //White Team Color
	blackColor := color.RGBA{0x00, 0x00, 0x00, 0xFF} //Black Team Color

	for i:=0;i<64;i++{
		centerX, centerY := lichessDimensions.getTilePos(i,0.5,0.5)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Empty { continue }
		
		if tile.Team == board.White {
			imd.Color = whiteColor
		} else {
			imd.Color = blackColor
		}
		imd.Push(pixel.V(centerX, centerY))
	}

	imd.Ellipse(pixel.V(size/2.5, size/2.5), 0.0)

	for i:=0;i<64;i++ {
		centerX, centerY := lichessDimensions.getTilePos(i,0.5,0.5)

		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Empty || tile.King == board.Pawn { continue }
		
		if tile.Team == board.White {
			imd.Color = blackColor
		} else {
			imd.Color = whiteColor
		}
		imd.Push(pixel.V(centerX, centerY))
	}

	imd.Ellipse(pixel.V(size/5.0, size/5.0),0.0)
}

func (t LichessTheme) DrawSelected(imd *imdraw.IMDraw, index int) {
	if index == -1 { return }
	tileX, tileY := lichessDimensions.getTilePosBL(index)
	imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Selection Color
	imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + lichessDimensions.getTileSize(), tileY + lichessDimensions.getTileSize()))
	imd.Rectangle(0.0)
}

func (t LichessTheme) DrawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	if index == -1 { return }
	moves, exist := moveMap[index]
	if exist {
		imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Move Corner Color
		for _, imove := range moves {
			corners(imd, imove, 20.0, lichessDimensions)
		}
	}
}

func (t LichessTheme) DrawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	imd.Color = color.RGBA{0xFF, 0x00, 0x00, 0x7F} //Take Corner Color
	for a := range moveMap {
		corners(imd, a, 20.0, lichessDimensions)
	}
}

