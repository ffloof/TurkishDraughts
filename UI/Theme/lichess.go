package theme

import (
	"TurkishDraughts/Board"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)



type LichessTheme struct {} 



var (
	
	lichessBoardBg = color.RGBA{0xED, 0xEB, 0xE9, 0xFF}
	lichessTileDark = color.RGBA{0xB5, 0x88, 0x63, 0xFF}
	lichessTileLight = color.RGBA{0xF0, 0xD9, 0xB5, 0xFF}
	
	lichessTeamWhite = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	lichessTeamBlack = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	lichessTakeColor = color.RGBA{0xFF, 0x00, 0x00, 0x7F}
	lichessMoveColor = color.RGBA{0x00, 0x3F, 0x00, 0x7F}

	lichessCornerSize = 20.0

	lichessHoverSize = 2.0
)

var lichessDimensions = dimensions {
	Height: 900.0,
	Border: 20.0,
	Gaps: 4.0,
	InternalGap: 4.0,
}


//TODO: move to utils.go
func (l LichessTheme) GetMouseData(win *pixelgl.Window) (bool, bool, int) {
	mPos := win.MousePosition()
	if mPos.X > lichessDimensions.Height - lichessDimensions.Border - lichessDimensions.Gaps { return false, false, -1 }
	if mPos.Y > lichessDimensions.Height - lichessDimensions.Border - lichessDimensions.Gaps { return false, false, -1 }

	tileX := int((mPos.X - lichessDimensions.Border) / lichessDimensions.getTileSpace())
	tileY := 7 - int((mPos.Y - lichessDimensions.Border) / lichessDimensions.getTileSpace())
	
	return win.JustPressed(pixelgl.MouseButtonLeft), win.JustReleased(pixelgl.MouseButtonLeft), (tileY * 8) + tileX
}

func (l LichessTheme) DrawBoard(imd *imdraw.IMDraw){
	imd.Color = lichessBoardBg
	imd.Push(pixel.V(0.0,0.0), pixel.V(lichessDimensions.Height, lichessDimensions.Height))
	imd.Rectangle(0.0)

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if (x+y)%2 == 0 {
				imd.Color = lichessTileDark
			} else {
				imd.Color = lichessTileLight
			}

			x1, y1 := lichessDimensions.getTilePos(x,y,0.0,0.0)
			x2, y2 := lichessDimensions.getTilePos(x,y,1.0,1.0)

			imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
			imd.Rectangle(0.0)
		}
	}
}

func (l LichessTheme) DrawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	size := lichessDimensions.getTileSpace()
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			centerX, centerY := lichessDimensions.getTilePos(x,y,0.5,0.5)

			tile, _ := b.GetBoardTile(x,y)
			if tile.Full == board.Empty { continue }
			
			if tile.Team == board.White {
				imd.Color = lichessTeamWhite
			} else {
				imd.Color = lichessTeamBlack
			}
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V((size/2.0)-(lichessDimensions.InternalGap)-(lichessDimensions.Gaps), (size/2.0)-(lichessDimensions.InternalGap)-(lichessDimensions.Gaps)),0.0)

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			centerX, centerY := lichessDimensions.getTilePos(x,y,0.5,0.5)

			tile, _ := b.GetBoardTile(x,y)
			if tile.Full == board.Empty || tile.King == board.Pawn { continue }
			
			if tile.Team == board.White {
				imd.Color = lichessTeamBlack
			} else {
				imd.Color = lichessTeamWhite
			}
			imd.Push(pixel.V(centerX, centerY))
		}
	}
	imd.Ellipse(pixel.V((size/4.0)-(lichessDimensions.InternalGap)-(lichessDimensions.Gaps), (size/4.0)-(lichessDimensions.InternalGap)-(lichessDimensions.Gaps)),0.0)
}

func (l LichessTheme) DrawSelected(imd *imdraw.IMDraw, index int) {
	if index == -1 { return }
	tileX, tileY := lichessDimensions.getTilePosBL(index%8, index/8)
	imd.Color = lichessMoveColor
	imd.Push(pixel.V(tileX, tileY), pixel.V(tileX + lichessDimensions.getTileSize(), tileY + lichessDimensions.getTileSize()))
	imd.Rectangle(0.0)
}

func (l LichessTheme) DrawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	if index == -1 { return }
	moves, exist := moveMap[index]
	if exist {
		imd.Color = lichessMoveColor
		for _, imove := range moves {
			corners(imd, imove%8, imove/8, lichessCornerSize, lichessDimensions)
		}
	}
}

func (l LichessTheme) DrawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	imd.Color = lichessTakeColor
	for a := range moveMap {
		corners(imd, a%8, a/8, lichessCornerSize, lichessDimensions)
	}
}

