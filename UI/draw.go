package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)

var (
	Border = 20.0
	Gaps = 4.0
	
	BoardBg = color.RGBA{0xED, 0xEB, 0xE9, 0xFF}
	TileDark = color.RGBA{0xB5, 0x88, 0x63, 0xFF}
	TileLight = color.RGBA{0xF0, 0xD9, 0xB5, 0xFF}
	
	TeamWhite = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	TeamBlack = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	HoverColor = color.RGBA{0x00, 0x88, 0x00, 0xFF}
	TakeColor = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	MoveColor = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	InternalGap = 4.0
)

func getTileSpace() float64 {
	return (float64(Height) - (2*Border) + Gaps)/ 8.0
}

func getTileSize() float64{
	return (getTileSpace() - Gaps)
}

func getTilePosBL(x int, y int) (float64, float64) {
	return float64(x)*getTileSpace() + Border, float64(y)*getTileSpace() + Border
	
}

func getTilePosCenter(x int, y int) (float64, float64) {
	return (float64(x)+0.5)*getTileSpace() + Border, (float64(y)+0.5)*getTileSpace() + Border
}


func drawBoard(imd *imdraw.IMDraw){
	imd.Color = BoardBg
	imd.Push(pixel.V(0.0,0.0), pixel.V(float64(Height), float64(Height)))
	imd.Rectangle(0.0)

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX, posY := getTilePosBL(x,y)
			
			if (x+y)%2 == 0 {
				imd.Color = TileDark
			} else {
				imd.Color = TileLight
			}
			imd.Push(pixel.V(posX, posY), pixel.V(posX+getTileSize(), posY+getTileSize()))
			imd.Rectangle(0.0)
		}
	}
}

func drawPieces(b *board.BoardState, imd *imdraw.IMDraw){
	size := getTileSpace()
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX, posY := getTilePosCenter(x,y)

			tile, _ := b.GetBoardTile(x,7-y)
			if tile.Full == board.Empty { continue }
			
			if tile.Team == board.White {
				imd.Color = TeamWhite
			} else {
				imd.Color = TeamBlack
			}
			imd.Push(pixel.V(posX-(Gaps/2.0), posY-(Gaps/2.0)))
		}
	}
	imd.Ellipse(pixel.V((size/2.0)-(InternalGap)-(Gaps), (size/2.0)-(InternalGap)-(Gaps)),0.0)

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX, posY := getTilePosCenter(x,y)

			tile, _ := b.GetBoardTile(x,7-y)
			if tile.Full == board.Empty || tile.King == board.Pawn { continue }
			
			if tile.Team == board.White {
				imd.Color = TeamBlack
			} else {
				imd.Color = TeamWhite
			}
			imd.Push(pixel.V(posX-(Gaps/2.0), posY-(Gaps/2.0)))
		}
	}
	imd.Ellipse(pixel.V((size/4.0)-(InternalGap)-(Gaps), (size/4.0)-(InternalGap)-(Gaps)),0.0)
}

func drawControls(){

}

func drawHover(win *pixelgl.Window, imd *imdraw.IMDraw) (bool, int) {
	imd.Color = HoverColor
	size := getTileSize()
	mPos := win.MousePosition()
	if mPos.X > Height - Border - Gaps { return false, 0 }
	if mPos.Y > Height - Border - Gaps { return false, 0 }

	tileX := int((mPos.X - Border) / getTileSpace())
	tileY := int((mPos.Y - Border) / getTileSpace())

	tmpX, tmpY := getTilePosBL(tileX, tileY)
	imd.Push(pixel.V(tmpX, tmpY), pixel.V(tmpX + size, tmpY + size))

	imd.Rectangle(5.0)
	
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		return true, (tileY * 8) + tileX
	} else {
		return false, 0
	}
}

func drawSelected(imd *imdraw.IMDraw) {
	
}

func drawChecks(imd *imdraw.IMDraw, tilemap map[int][]int) {
	size := getTileSize()
	for a := range tilemap {
		tileX, tileY := getTilePosBL(a%8, 7-(a/8))
		imd.Color = TakeColor
		imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + size, tileY + size))
		imd.Rectangle(5.0)
	}
}