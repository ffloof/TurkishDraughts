package theme

//Theme based off the chess board/pieces at lichess.org, no affiliation.

import (
	"TurkishDraughts/Board"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)


type RainbowTheme struct {
} 

var rainbowDimensions = dimensions {
	Height: 900.0,
	Border: 20.0,
	Gaps: 0.0,
}

var Hue float64 = 0.0

func (r RainbowTheme) GetMouseData(win *pixelgl.Window) (bool, bool, int) {
	return rainbowDimensions.getMouseData(win)
}

func (r RainbowTheme) DrawBoard(imd *imdraw.IMDraw){
	imd.Color = color.RGBA{0xED, 0xEB, 0xE9, 0xFF} //Board background color
	imd.Push(pixel.V(0.0,0.0), pixel.V(rainbowDimensions.Height, rainbowDimensions.Height))
	imd.Rectangle(0.0)

	Hue += 0.001
	if Hue > 1.0 { Hue = 0.0 }

	for i:=0;i<64;i++ {
		if (i/8 + i%8)%2 == 0 { //Checkerboard effect (x+y)%2
			imd.Color = HSLToRGB(Hue, 1.0, 0.666) //Tile Dark Color
		} else {
			imd.Color = HSLToRGB(Hue, 1.0, 0.333) //Tile Light Color
		}
		x1, y1 := rainbowDimensions.getTilePos(i,0.0,0.0)
		x2, y2 := rainbowDimensions.getTilePos(i,1.0,1.0)
		imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
		imd.Rectangle(0.0)
	}
}




func (r RainbowTheme) DrawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	size := rainbowDimensions.getTileSpace()

	whiteColor := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF} //White Team Color
	blackColor := color.RGBA{0x00, 0x00, 0x00, 0xFF} //Black Team Color

	for i:=0;i<64;i++{
		centerX, centerY := rainbowDimensions.getTilePos(i,0.5,0.5)
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
		centerX, centerY := rainbowDimensions.getTilePos(i,0.5,0.5)

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

func (r RainbowTheme) DrawSelected(imd *imdraw.IMDraw, index int) {
	tileX, tileY := rainbowDimensions.getTilePosBL(index)
	imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Selection Color
	imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + rainbowDimensions.getTileSize(), tileY + rainbowDimensions.getTileSize()))
	imd.Rectangle(0.0)
}

func (r RainbowTheme) DrawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	moves, exist := moveMap[index]
	if exist {
		imd.Color = color.RGBA{0x00, 0x3F, 0x00, 0x7F} //Move Corner Color
		for _, imove := range moves {
			corners(imd, imove, 20.0, rainbowDimensions)
		}
	}
}

func (r RainbowTheme) DrawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	imd.Color = color.RGBA{0xFF, 0x00, 0x00, 0x7F} //Take Corner Color
	for a := range moveMap {
		corners(imd, a, 20.0, rainbowDimensions)
	}
}

