package theme

//Theme based off the chess board/pieces on the wikipedia page for Turkish Draughts, no affiliation.

import (
	"TurkishDraughts/Board"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)


type WikipediaTheme struct {} 

var wikipediaDimensions = dimensions {
	Height: 900.0,
	Border: 3.0,
	Gaps: 3.0,
}

func (t WikipediaTheme) GetMouseData(win *pixelgl.Window) (bool, bool, int) {
	return wikipediaDimensions.getMouseData(win)
}

func (t WikipediaTheme) DrawBoard(imd *imdraw.IMDraw){
	imd.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF} //Board background color
	imd.Push(pixel.V(0.0,0.0), pixel.V(wikipediaDimensions.Height, wikipediaDimensions.Height))
	imd.Rectangle(0.0)

	for i:=0;i<64;i++ {
		imd.Color = color.RGBA{0xFF, 0xEE, 0xBB, 0xFF} //Board Tile Background Color
		x1, y1 := wikipediaDimensions.getTilePos(i,0.0,0.0)
		x2, y2 := wikipediaDimensions.getTilePos(i,1.0,1.0)
		imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
		imd.Rectangle(0.0)
	}
}

func (t WikipediaTheme) DrawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	size := wikipediaDimensions.getTileSpace()

	outlineThickness := 2.0

	imd.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	for i:=0;i<64;i++{
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.45)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled {
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V(size/3.0+outlineThickness, size/4.0+outlineThickness),0.0)

	for i:=0;i<64;i++{
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.5)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled {
			if tile.Team == board.White {
				imd.Color = color.RGBA{0xFF, 0xF9, 0xF4, 0xFF} //White Team Color
			} else {
				imd.Color = color.RGBA{0xC4, 0x00, 0x03, 0xFF} //Black/Red Team Color
			}
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V(size/3.0, size/4.0),0.0)

	imd.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	for i:=0;i<64;i++{
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.5)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled {
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V(size/3.0, size/4.0),outlineThickness*2.0)

	imd.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	for i:=0;i<64;i++{
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.5)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled && tile.King == board.King{
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V(size/9.0, size/12.0),0.0)

	/*
	for i:=0;i<64;i++{
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.4)
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Empty { continue }
		
		if tile.Team == board.White {
			imd.Color = whiteColor
		} else {
			imd.Color = blackColor
		}
		imd.Push(pixel.V(centerX, centerY))
	}

	imd.Ellipse(pixel.V(size/2.5, size/2.5),0.0)

	for i:=0;i<64;i++ {
		centerX, centerY := wikipediaDimensions.getTilePos(i,0.5,0.5)

		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Empty || tile.King == board.Pawn { continue }
		
		if tile.Team == board.White {
			imd.Color = blackColor
		} else {
			imd.Color = whiteColor
		}
		imd.Push(pixel.V(centerX, centerY))
	}

	imd.Ellipse(pixel.V(size/4.0, size/4.0),0.0)*/
}

func (t WikipediaTheme) DrawSelected(imd *imdraw.IMDraw, index int) {
	if index == -1 { return }
	tileX, tileY := wikipediaDimensions.getTilePosBL(index)
	imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Selection Color
	imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + wikipediaDimensions.getTileSize(), tileY + wikipediaDimensions.getTileSize()))
	imd.Rectangle(0.0)
}

func (t WikipediaTheme) DrawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	if index == -1 { return }
	moves, exist := moveMap[index]
	if exist {
		imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Move Corner Color
		for _, imove := range moves {
			corners(imd, imove, 20.0, wikipediaDimensions)
		}
	}
}

func (t WikipediaTheme) DrawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	imd.Color = color.RGBA{0xFF, 0x00, 0x00, 0x7F} //Take Corner Color
	for a := range moveMap {
		corners(imd, a, 20.0, wikipediaDimensions)
	}
}