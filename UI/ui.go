package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Width = 1600
	Height = 900
)

var (
	Border = 20.0
	Gaps = 4.0
	BoardBg = color.RGBA{0xED, 0xEB, 0xE9, 0xFF}
	TileA = color.RGBA{0xB5, 0x88, 0x63, 0xFF}
	TileB = color.RGBA{0xF0, 0xD9, 0xB5, 0xFF}

	InternalGap = 10.0
)


func Init() {
	b := board.CreateStartingBoard()

	cfg := pixelgl.WindowConfig{
		Title:  "Turkish Draughts Engine",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	/*
	imd.Color = colornames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.EndShape = imdraw.SharpEndShape
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(30)

	imd.Color = colornames.Limegreen
	imd.Push(pixel.V(500, 500))
	imd.Circle(300, 50)

	imd.Color = colornames.Navy
	imd.Push(pixel.V(200, 500), pixel.V(800, 500))
	imd.Ellipse(pixel.V(120, 80), 0)

	imd.Color = colornames.Red
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(500, 350))
	imd.CircleArc(150, -math.Pi, 0, 30)
	*/

	for !win.Closed() {
		imd := imdraw.New(nil)
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
		
		drawBoard(imd)
		drawPieces(&b, imd)

		imd.Draw(win)
		win.Update()
	}
}

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
				imd.Color = TileA
			} else {
				imd.Color = TileB
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

			tile, _ := b.GetBoardTile(x,y)
			if tile.Full == board.Empty { continue }
			
			if tile.Team == board.White {
				imd.Color = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
			} else {
				imd.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
			}
			imd.Push(pixel.V(posX-(Gaps/2.0), posY-(Gaps/2.0)))
		}
	}
	imd.Ellipse(pixel.V((size/2.0)-(InternalGap)-(Gaps), (size/2.0)-(InternalGap)-(Gaps)),0.0)
}

func drawControls(){

}