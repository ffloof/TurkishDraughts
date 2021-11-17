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

	TakeColor = color.RGBA{0xFF, 0x00, 0x00, 0x7F}
	MoveColor = color.RGBA{0x00, 0x3F, 0x00, 0x7F}

	HairLength = 20.0
	HairSize = 5.0

	HoverSize = 2.0

	InternalGap = 4.0
)

func getTileSpace() float64 {
	return (float64(Height) - (2*Border) + Gaps)/ 8.0
}

func getTileSize() float64{
	return (getTileSpace() - Gaps)
}

func getTilePosBL(x int, y int) (float64, float64) {
	return float64(x)*getTileSpace() + Border, float64(7-y)*getTileSpace() + Border
	
}

func getTilePosCenter(x int, y int) (float64, float64) {
	return (float64(x)+0.5)*getTileSpace() + Border, (float64(7-y)+0.5)*getTileSpace() + Border
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

func drawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	size := getTileSpace()
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			posX, posY := getTilePosCenter(x,y)

			tile, _ := b.GetBoardTile(x,y)
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

			tile, _ := b.GetBoardTile(x,y)
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

func getMouseData(win *pixelgl.Window) (bool, bool, int) {
	mPos := win.MousePosition()
	if mPos.X > Height - Border - Gaps { return false, false, -1 }
	if mPos.Y > Height - Border - Gaps { return false, false, -1 }

	tileX := int((mPos.X - Border) / getTileSpace())
	tileY := 7 - int((mPos.Y - Border) / getTileSpace())
	
	return win.JustPressed(pixelgl.MouseButtonLeft), win.JustReleased(pixelgl.MouseButtonLeft), (tileY * 8) + tileX
}

func drawSelected(imd *imdraw.IMDraw, index int) {
	if index == -1 { return }
	tileX, tileY := getTilePosBL(index%8, index/8)
	imd.Color = MoveColor
	imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + getTileSize(), tileY + getTileSize()))
	imd.Rectangle(0.0)
}

func drawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	if index == -1 { return }
	moves, exist := moveMap[index]
	if exist {
		for _, imove := range moves {
			tileX, tileY := getTilePosBL(imove%8, imove/8)
			imd.Color = MoveColor
			corners(imd, tileX, tileY)
		}
	}
}

func drawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	for a := range moveMap {
		tileX, tileY := getTilePosBL(a%8, a/8)
		imd.Color = TakeColor
		corners(imd, tileX, tileY)
	}
}

func corners(imd *imdraw.IMDraw, x1, y1 float64){
	size := getTileSize()
	x2 := x1 + size
	y2 := y1 + size

	//Bottom left corner
	imd.Push(pixel.V(x1+HairLength,y1), pixel.V(x1,y1), pixel.V(x1,y1+HairLength))
	imd.Polygon(0)

	//Bottom right corner
	imd.Push(pixel.V(x2-HairLength,y1), pixel.V(x2,y1), pixel.V(x2,y1+HairLength))
	imd.Polygon(0)

	//Top left corner
	imd.Push(pixel.V(x1+HairLength,y2), pixel.V(x1,y2), pixel.V(x1,y2-HairLength))
	imd.Polygon(0)

	//Top right corner
	imd.Push(pixel.V(x2-HairLength,y2), pixel.V(x2,y2), pixel.V(x2,y2-HairLength))
	imd.Polygon(0)
}