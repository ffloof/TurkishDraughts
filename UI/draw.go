package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"github.com/faiface/pixel"
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

	InternalGap = 4.0
)

func drawBoard(imd *imdraw.IMDraw){
	imd.Color = BoardBg
	imd.Push(pixel.V(0.0,0.0), pixel.V(float64(Height), float64(Height)))
	imd.Rectangle(0.0)

	size := (float64(Height) - (2*Border) + Gaps)/ 8.0
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX := float64(x)*size + Border
			posY := float64(y)*size + Border
			if (x+y)%2 == 0 {
				imd.Color = TileDark
			} else {
				imd.Color = TileLight
			}
			imd.Push(pixel.V(posX, posY), pixel.V(posX+size-Gaps, posY+size-Gaps))
			imd.Rectangle(0.0)
		}
	}
}

func drawPieces(b *board.BoardState, imd *imdraw.IMDraw){
	size := (float64(Height) - (2*Border) + Gaps)/ 8.0
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX := (float64(x)+0.5)*size + Border
			posY := (float64(y)+0.5)*size + Border

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
			posX := (float64(x)+0.5)*size + Border
			posY := (float64(y)+0.5)*size + Border

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